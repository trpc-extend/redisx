package redisx

import (
	"fmt"
	"net"
	"testing"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/smartystreets/goconvey/convey"

	"git.code.oa.com/goom/mocker"
	"git.code.oa.com/trpc-go/trpc-database/redis"
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/codec"
	"git.code.oa.com/trpc-go/trpc-go/transport"
)

// TestNewClientTransport
func TestNewClientTransport(t *testing.T) {
	convey.Convey("succ", t, func() {
		mock := mocker.Create()
		defer mock.Reset()

		opts := []transport.ClientTransportOption{transport.WithClientUDPRecvSize(1)}

		ret := NewClientTransport(opts...)
		convey.So(ret, convey.ShouldNotBeNil)
	})
}

// TestHasQuota
func TestRoundTrip_Failed(t *testing.T) {
	ctx := trpc.BackgroundContext()
	_, msg := codec.WithNewMessage(ctx)

	callOpts := []transport.RoundTripOption{transport.WithDialAddress("1")}

	req := &redis.Request{
		Command:  "GET",
		Args:     []interface{}{"k1", "k2"},
		Pipeline: false,
	}

	rsp := &redis.Response{
		Reply: "",
	}

	ct := &ClientTransport{
		opts:            &transport.ClientTransportOptions{},
		redisPool:       make(map[string]*redigo.Pool),
		MaxIdle:         2048,
		MaxActive:       0,
		IdleTimeout:     3 * time.Minute,
		MaxConnLifetime: 0,
		DefaultTimeout:  time.Second,
	}

	conn := ct.GetConn("2", "2", time.Duration(0))

	convey.Convey("redis.Request err", t, func() {
		mock := mocker.Create()
		defer mock.Reset()

		mock.Struct(msg).Method("ClientReqHead").Return(nil)

		rsp, err := ct.RoundTrip(ctx, nil, callOpts...)
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(rsp, convey.ShouldBeNil)
	})

	convey.Convey("redis.Response err", t, func() {
		mock := mocker.Create()
		defer mock.Reset()

		mock.Struct(msg).Method("ClientReqHead").Return(req)
		mock.Struct(msg).Method("ClientRspHead").Return(nil)

		rsp, err := ct.RoundTrip(ctx, nil, callOpts...)
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(rsp, convey.ShouldBeNil)
	})

	convey.Convey("DoWithTimeout err", t, func() {
		mock := mocker.Create()
		defer mock.Reset()

		req.Pipeline = false

		mock.Struct(msg).Method("ClientReqHead").Return(req)
		mock.Struct(msg).Method("ClientRspHead").Return(rsp)
		mock.Struct(ct).Method("GetConn").Return(conn)

		netErr := &net.OpError{
			Err: fmt.Errorf("connection refused"),
		}
		mock.Func(redigo.DoWithTimeout).Return(nil, netErr)

		rsp, err := ct.RoundTrip(ctx, nil, callOpts...)
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(rsp, convey.ShouldBeNil)

		mock.Func(redigo.DoWithTimeout).Return(nil, fmt.Errorf("not timeout"))
		rsp, err = ct.RoundTrip(ctx, nil, callOpts...)
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(rsp, convey.ShouldBeNil)
	})
}

func TestRoundTrip_Succ(t *testing.T) {
	ctx := trpc.BackgroundContext()
	_, msg := codec.WithNewMessage(ctx)

	callOpts := []transport.RoundTripOption{transport.WithDialAddress("1")}

	req := &redis.Request{
		Command:  "GET",
		Args:     []interface{}{"k1", "k2"},
		Pipeline: false,
	}

	rsp := &redis.Response{
		Reply: "",
	}

	ct := &ClientTransport{
		opts:            &transport.ClientTransportOptions{},
		redisPool:       make(map[string]*redigo.Pool),
		MaxIdle:         2048,
		MaxActive:       0,
		IdleTimeout:     3 * time.Minute,
		MaxConnLifetime: 0,
		DefaultTimeout:  time.Second,
	}

	conn := ct.GetConn("2", "2", time.Duration(0))

	convey.Convey("Pipeline true", t, func() {
		mock := mocker.Create()
		defer mock.Reset()

		req.Pipeline = true
		mock.Struct(msg).Method("ClientReqHead").Return(req)
		mock.Struct(msg).Method("ClientRspHead").Return(rsp)
		mock.Struct(ct).Method("GetConn").Return(conn)

		rsp, err := ct.RoundTrip(ctx, nil, callOpts...)
		convey.So(err, convey.ShouldBeNil)
		convey.So(rsp, convey.ShouldBeNil)
	})

	convey.Convey("succ", t, func() {
		mock := mocker.Create()
		defer mock.Reset()

		req.Pipeline = false
		mock.Struct(msg).Method("ClientReqHead").Return(req)
		mock.Struct(msg).Method("ClientRspHead").Return(rsp)
		mock.Struct(ct).Method("GetConn").Return(conn)
		mock.Func(redigo.DoWithTimeout).Return(nil, nil)

		rsp, err := ct.RoundTrip(ctx, nil, callOpts...)
		convey.So(err, convey.ShouldBeNil)
		convey.So(rsp, convey.ShouldBeNil)
	})
}

func TestGetConn(t *testing.T) {
	connPool := map[string]*redigo.Pool{}
	pool := &redigo.Pool{
		MaxIdle:         0,
		MaxActive:       0,
		IdleTimeout:     0,
		MaxConnLifetime: 0,
		Dial: func() (redigo.Conn, error) {
			return redigo.DialURL(
				"redis://addr",
				redigo.DialConnectTimeout(time.Duration(0)),
				redigo.DialWriteTimeout(time.Duration(0)),
				redigo.DialReadTimeout(time.Duration(0)),
				redigo.DialPassword("pwd"),
			)
		},
	}
	connPool["1:1:0"] = pool

	ct := &ClientTransport{
		opts:            &transport.ClientTransportOptions{},
		redisPool:       connPool,
		MaxIdle:         2048,
		MaxActive:       0,
		IdleTimeout:     3 * time.Minute,
		MaxConnLifetime: 0,
		DefaultTimeout:  time.Second,
	}

	convey.Convey("pool geted", t, func() {
		mock := mocker.Create()
		defer mock.Reset()

		ret := ct.GetConn("1", "1", time.Duration(0))
		convey.So(ret, convey.ShouldNotBeNil)
	})

	convey.Convey("new pool", t, func() {
		mock := mocker.Create()
		defer mock.Reset()

		ret := ct.GetConn("2", "2", time.Duration(0))
		convey.So(ret, convey.ShouldNotBeNil)
	})

}
