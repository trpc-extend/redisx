package redisx

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	redigo "github.com/gomodule/redigo/redis"

	"git.code.oa.com/trpc-go/trpc-database/redis"
	"git.code.oa.com/trpc-go/trpc-go/codec"
	"git.code.oa.com/trpc-go/trpc-go/errs"
	"git.code.oa.com/trpc-go/trpc-go/naming/selector"
	"git.code.oa.com/trpc-go/trpc-go/transport"
	dsn "git.code.oa.com/trpc-go/trpc-selector-dsn"
)

func init() {
	transport.RegisterClientTransport("redis", DefaultClientTransport)
	selector.Register("redis+polaris", dsn.NewResolvableSelector("polaris", &dsn.URIHostExtractor{}))
	selector.Register("redis", dsn.DefaultSelector)
}

// ClientTransport client端redis transport
type ClientTransport struct {
	opts            *transport.ClientTransportOptions
	redisPool       map[string]*redigo.Pool
	redisPoolLock   sync.RWMutex
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	MaxConnLifetime time.Duration
	DefaultTimeout  time.Duration // 设置默认连接超时时间,1s
}

// DefaultClientTransport 默认client redis transport
var DefaultClientTransport = NewClientTransport()

// NewClientTransport 创建redis transport
func NewClientTransport(opt ...transport.ClientTransportOption) transport.ClientTransport {

	opts := &transport.ClientTransportOptions{}

	// 将传入的func option写到opts字段中
	for _, o := range opt {
		o(opts)
	}

	return &ClientTransport{
		opts:            opts,
		redisPool:       make(map[string]*redigo.Pool),
		MaxIdle:         2048,
		MaxActive:       0,
		IdleTimeout:     3 * time.Minute,
		MaxConnLifetime: 0,
		DefaultTimeout:  time.Second,
	}
}

// RoundTrip 收发redis包, 回包redis response放到ctx里面，这里不需要返回rspbuf
func (ct *ClientTransport) RoundTrip(ctx context.Context, _ []byte,
	callOpts ...transport.RoundTripOption) (rspByte []byte, err error) {

	msg := codec.Message(ctx)

	req, ok := msg.ClientReqHead().(*redis.Request)
	if !ok {
		return nil, errs.NewFrameError(errs.RetClientEncodeFail,
			"redis client transport: ReqHead should be type of *redis.Request")
	}
	rsp, ok := msg.ClientRspHead().(*redis.Response)
	if !ok {
		return nil, errs.NewFrameError(errs.RetClientEncodeFail,
			"redis client transport: RspHead should be type of *redis.Response")
	}

	opts := &transport.RoundTripOptions{}
	for _, o := range callOpts {
		o(opts)
	}

	conn := ct.GetConn(opts.Address, opts.Password, msg.RequestTimeout())
	if req.Pipeline {
		rsp.Reply = conn
		return nil, nil
	}
	defer conn.Close()

	reply, err := redigo.DoWithTimeout(conn, msg.RequestTimeout(), req.Command, req.Args...)
	if err != nil {
		if e, ok := err.(net.Error); ok {
			if e.Timeout() {
				return nil, errs.NewFrameError(errs.RetClientTimeout, err.Error())
			}
			if strings.Contains(err.Error(), "connection refused") {
				return nil, errs.NewFrameError(errs.RetClientConnectFail, err.Error())
			}
		}
		return nil, errs.NewFrameError(errs.RetClientNetErr, err.Error())
	}
	rsp.Reply = reply

	return nil, nil
}

// GetConn 获取redis链接
func (ct *ClientTransport) GetConn(address string, password string, timeout time.Duration) redigo.Conn {
	key := fmt.Sprintf("%s:%s:%d", address, password, FormatMs(timeout))

	ct.redisPoolLock.RLock()
	pool, ok := ct.redisPool[key]
	ct.redisPoolLock.RUnlock()
	if ok {
		return pool.Get()
	}

	ct.redisPoolLock.Lock()
	defer ct.redisPoolLock.Unlock()
	pool, ok = ct.redisPool[key]
	if ok {
		return pool.Get()
	}

	//超时时间没有配置,则使用默认的超时时间(1s)
	dialTimeout := timeout
	if FormatMs(dialTimeout) == 0 {
		dialTimeout = ct.DefaultTimeout
	}

	pool = &redigo.Pool{
		MaxIdle:         ct.MaxIdle,
		MaxActive:       ct.MaxActive,
		IdleTimeout:     ct.IdleTimeout,
		MaxConnLifetime: ct.MaxConnLifetime,
		Dial: func() (redigo.Conn, error) {
			return redigo.DialURL(
				"redis://"+address,
				redigo.DialConnectTimeout(dialTimeout),
				redigo.DialWriteTimeout(dialTimeout),
				redigo.DialReadTimeout(dialTimeout),
				redigo.DialPassword(password),
			)
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	ct.redisPool[key] = pool

	return pool.Get()
}
