package redisx

import (
	"context"
	"time"

	redigo "github.com/gomodule/redigo/redis"

	"git.code.oa.com/trpc-go/trpc-database/redis"
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/log"
)

// Client redis请求接口
type Client interface {
	Do(ctx context.Context, cmd string, args ...interface{}) (interface{}, error)
	RetryBackoff(retry int, minBackoff, maxBackoff time.Duration)

	// pipeline
	Pipeline(ctx context.Context, request *PipelineRequest) (*PipelineResponse, error)
	GetPipelineConn(ctx context.Context) (redigo.Conn, error)

	// script
	Script(ctx context.Context, keyCount int, script string, keysAndArgs ...interface{}) (interface{}, error)

	// strings
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	PSetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	SetRange(ctx context.Context, key string, offset int64, value string) (int64, error)
	Get(ctx context.Context, key string) (string, error)
	GetRange(ctx context.Context, key string, start, end int64) (string, error)
	GetSet(ctx context.Context, key string, value interface{}) (string, error)
	Incr(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
	IncrByFloat(ctx context.Context, key string, value float64) (float64, error)
	Decr(ctx context.Context, key string) (int64, error)
	DecrBy(ctx context.Context, key string, decrement int64) (int64, error)
	StrLen(ctx context.Context, key string) (int64, error)
	Append(ctx context.Context, key, value string) (int64, error)
	MGet(ctx context.Context, keys ...string) ([]string, error)
	MSet(ctx context.Context, values ...interface{}) error
	MSetNX(ctx context.Context, values ...interface{}) (bool, error)

	// list
	BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error)
	LIndex(ctx context.Context, key string, index int64) (string, error)
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error)
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error)
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error)
	LLen(ctx context.Context, key string) (int64, error)
	LPop(ctx context.Context, key string) (string, error)
	LPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	LPushX(ctx context.Context, key string, values ...interface{}) (int64, error)
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error)
	LSet(ctx context.Context, key string, index int64, value interface{}) error
	LTrim(ctx context.Context, key string, start, stop int64) error
	RPop(ctx context.Context, key string) (string, error)
	RPopLPush(ctx context.Context, source, destination string) (string, error)
	RPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	RPushX(ctx context.Context, key string, value interface{}) (int64, error)

	// set
	SAdd(ctx context.Context, key string, members ...interface{}) (int64, error)
	SCard(ctx context.Context, key string) (int64, error)
	SDiff(ctx context.Context, keys ...string) ([]string, error)
	SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error)
	SInter(ctx context.Context, keys ...string) ([]string, error)
	SInterStore(ctx context.Context, destination string, keys ...string) (int64, error)
	SIsMember(ctx context.Context, key string, member interface{}) (bool, error)
	SMembers(ctx context.Context, key string) ([]string, error)
	SMove(ctx context.Context, source, destination string, member interface{}) (bool, error)
	SPop(ctx context.Context, key string) (string, error)
	SPopN(ctx context.Context, key string, count int64) ([]string, error)
	SRandMember(ctx context.Context, key string) (string, error)
	SRandMemberN(ctx context.Context, key string, count int64) ([]string, error)
	SRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	SUnion(ctx context.Context, keys ...string) ([]string, error)
	SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error)

	// keys
	Del(ctx context.Context, keys ...string) (int64, error)
	Exists(ctx context.Context, keys ...string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error)
	Keys(ctx context.Context, pattern string) ([]string, error)
	Persist(ctx context.Context, key string) (bool, error)
	PExpire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	PExpireAt(ctx context.Context, key string, tm time.Time) (bool, error)
	PTTL(ctx context.Context, key string) (time.Duration, error)
	RandomKey(ctx context.Context) (string, error)
	Touch(ctx context.Context, keys ...string) (int64, error)
	TTL(ctx context.Context, key string) (time.Duration, error)
	Type(ctx context.Context, key string) (string, error)

	// hash
	HSet(ctx context.Context, key string, values ...interface{}) (int64, error)
	HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error)
	HGet(ctx context.Context, key, field string) (string, error)
	HExists(ctx context.Context, key, field string) (bool, error)
	HDel(ctx context.Context, key string, fields ...string) (int64, error)
	HLen(ctx context.Context, key string) (int64, error)
	HStrLen(ctx context.Context, key string, field string) (int64, error)
	HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error)
	HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error)
	HMSet(ctx context.Context, key string, values ...interface{}) (bool, error)
	HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	HKeys(ctx context.Context, key string) ([]string, error)
	HVals(ctx context.Context, key string) ([]string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)

	// zset
	ZAdd(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddNX(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddXX(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddCh(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddNXCh(ctx context.Context, key string, members ...*Z) (int64, error)
	ZAddXXCh(ctx context.Context, key string, members ...*Z) (int64, error)
	ZIncr(ctx context.Context, key string, member Z) (float64, error)
	ZIncrNX(ctx context.Context, key string, member Z) (float64, error)
	ZIncrXX(ctx context.Context, key string, member Z) (float64, error)
	ZCard(ctx context.Context, key string) (int64, error)
	ZCount(ctx context.Context, key, min, max string) (int64, error)
	ZLexCount(ctx context.Context, key, min, max string) (int64, error)
	ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error)
	ZMScore(ctx context.Context, key string, members ...string) ([]interface{}, error)
	ZPopMax(ctx context.Context, key string, count ...int64) ([]Z, error)
	ZPopMin(ctx context.Context, key string, count ...int64) ([]Z, error)
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error)
	ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) ([]string, error)
	ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) ([]Z, error)
	ZRank(ctx context.Context, key, member string) (int64, error)
	ZRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error)
	ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error)
	ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error)
	ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) ([]string, error)
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) ([]Z, error)
	ZRevRank(ctx context.Context, key, member string) (int64, error)
	ZScore(ctx context.Context, key, member string) (float64, error)
	ZRandMember(ctx context.Context, key string, count int, withScores bool) ([]string, error)

	//CAS LOCK
	Lock(ctx context.Context, key string, expiration time.Duration) error
	UnLock(ctx context.Context)
}

// impDao 实现
type impDao struct {
	Client           redis.Client
	Option           client.Options
	ServiceName      string        //redis服务名
	Target           string        //redis BID
	CallerServerName string        //主调服务名
	MaxRetries       int           //最大重试次数,默认不重试
	MinRetryBackoff  time.Duration //重试间隔min ms, 间隔时间随机在[min,max]之间
	MaxRetryBackoff  time.Duration //重试间隔max ms
	CASKey           string        // cas lock key
	CASValue         int64         // cas lock value
}

// NewClient 新建一个redis后端请求代理 必传参数 redis服务名: trpc.redis.xxx.xxx
var NewClient = func(name string, opts ...client.Option) Client {
	imp := &impDao{
		Client:          redis.NewClientProxy(name, opts...),
		ServiceName:     name,
		MaxRetries:      0,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 16 * time.Millisecond,
	}

	for _, o := range opts {
		o(&imp.Option)
	}

	return imp
}

// PipelineCmd struct
type PipelineCmd struct {
	Cmd  string        //需要执行的命令
	Args []interface{} //命令参数
}

// PipelineCmdReply struct
type PipelineCmdReply struct {
	Err   error       //状态错误码
	Reply interface{} //返回值,类型不定,由使用方解析
}

// PipelineRequest pipeline请求
type PipelineRequest struct {
	Cmds        []PipelineCmd //命令集合
	IsErrIgnore bool          //单个命令失败是否忽略
}

// PipelineResponse pipeline响应
type PipelineResponse struct {
	CmdReply []PipelineCmdReply //命令执行结果集合
}

// RetryBackoff 设置全局接口重试次数和间隔
func (i *impDao) RetryBackoff(retry int, minBackoff, maxBackoff time.Duration) {
	i.MaxRetries = retry
	i.MinRetryBackoff = minBackoff
	i.MaxRetryBackoff = maxBackoff
}

// Do 命令执行
func (i *impDao) Do(ctx context.Context, cmd string, args ...interface{}) (interface{}, error) {
	var lastErr error
	for attempt := 0; attempt <= i.MaxRetries; attempt++ {

		nowTime := time.Now()
		val, err := i.process(ctx, attempt, cmd, args...)

		// redis服务名,可以是options中指定的target或者配置文件/options配置的serviceName
		if i.Target == "" {
			_ = i.Option.LoadClientConfig(i.ServiceName)
			i.Target = i.Option.Target
		}

		// 主调服务名,没有则使用传递的serviceName
		if i.CallerServerName == "" {
			i.CallerServerName = i.ServiceName

			cfg := trpc.GlobalConfig()
			if cfg.Server.App != "" && cfg.Server.Server != "" {
				i.CallerServerName = cfg.Server.App + "." + cfg.Server.Server
			}
		}

		cmdReport(i.Target, i.CallerServerName, cmd, nowTime, err)

		if err == nil || err == redis.ErrNil {
			return val, err
		}

		lastErr = err
	}

	return nil, lastErr
}

// GetPipelineConn 获取pipeline
func (i *impDao) GetPipelineConn(ctx context.Context) (redigo.Conn, error) {
	return i.Client.Pipeline(ctx)
}

// Pipeline  Pipeline接口,注意:req与rsp中命令响应序号是否一致,需要平台来保证
func (i *impDao) Pipeline(ctx context.Context, request *PipelineRequest) (response *PipelineResponse, err error) {
	// 处理完毕后上报数据
	nowTime := time.Now()
	defer func() {
		go i.pipelineReport(request, response, nowTime, err)
	}()

	if request == nil || len(request.Cmds) == 0 {
		return nil, ErrPipelineCmdsEmpty
	}

	conn, err := i.Client.Pipeline(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	totalCmdNum := 0 //send成功的命令
	// pipeline cmd send
	for _, cmder := range request.Cmds {
		err = conn.Send(cmder.Cmd, cmder.Args...)
		if err == nil {
			totalCmdNum++
		} else if err != nil && !request.IsErrIgnore {
			return nil, err
		}
	}

	// pipeline flush
	err = conn.Flush()
	if err != nil {
		return nil, err
	}

	// pipeline receiver
	response = &PipelineResponse{
		CmdReply: make([]PipelineCmdReply, totalCmdNum),
	}
	for i := 0; i < totalCmdNum; i++ {
		reply, err := conn.Receive()
		response.CmdReply[i] = PipelineCmdReply{
			Err:   err,
			Reply: reply,
		}
	}

	return response, nil
}

// Script 执行脚本命令
func (i *impDao) Script(ctx context.Context,
	keyCount int, script string, keysAndArgs ...interface{}) (val interface{}, err error) {
	nowTime := time.Now()
	defer func() {
		i.getService()
		cmdReport(i.Target, i.CallerServerName, "Script", nowTime, err)
	}()

	scriptCli := redis.NewScript(keyCount, script)
	return scriptCli.Do(ctx, i.Client, keysAndArgs...)
}

// process  底层命令执行(包含重试休眠)
func (i *impDao) process(ctx context.Context, attempt int, cmd string, args ...interface{}) (interface{}, error) {
	if attempt > 0 {
		err := Sleep(ctx, RetryBackoff(attempt, i.MinRetryBackoff, i.MaxRetryBackoff))
		if err != nil {
			return nil, err
		}
	}

	return i.Client.Do(ctx, cmd, args...)
}

// Lock cas加锁
func (i *impDao) Lock(ctx context.Context, key string, expiration time.Duration) (err error) {
	nowTime := time.Now()
	defer func() {
		i.getService()
		cmdReport(i.Target, i.CallerServerName, "Lock", nowTime, err)
	}()

	if key == "" {
		return ErrKeyEmpty
	}

	i.CASKey = key
	i.CASValue = time.Now().UnixNano() / 1e3
	bRet, err := i.SetNX(ctx, i.CASKey, i.CASValue, expiration)
	if !bRet || err != nil {
		return ErrCASLock
	}

	return nil
}

// UnLock cas解锁
func (i *impDao) UnLock(ctx context.Context) {
	var err error
	nowTime := time.Now()
	defer func() {
		i.getService()
		cmdReport(i.Target, i.CallerServerName, "UnLock", nowTime, err)
	}()

	if i.CASKey == "" {
		return
	}

	luaScript := `
	local delnum = 0
	local ret = redis.call("get", KEYS[1]) 
	if ret == ARGV[1] then 
		delnum = redis.call("del", KEYS[1]) 
	end 
	return tonumber(delnum)
	`

	script := redis.NewScript(1, luaScript)
	res, err := script.Do(ctx, i.Client, i.CASKey, i.CASValue)
	if res == nil || err != nil {
		log.ErrorContextf(ctx, "UnLock key [%s] error:%s", i.CASKey, err)
	}
}

// pipelineReport pipeLine上报
func (i *impDao) pipelineReport(request *PipelineRequest, response *PipelineResponse, nowTime time.Time, err error) {
	//加载服务配置
	i.getService()

	if err != nil || response == nil {
		// pipeline命令执行失败
		pipelineReport(i.Target, i.CallerServerName, nowTime, err, "", err, PipeLineCmdStatVal)
		return
	}

	// pipeline命令执行成功
	pipelienStatVal := PipeLineCmdStatVal
	rspLen := len(response.CmdReply) //rsp返回
	for idx, subCmd := range request.Cmds {
		cmdName := subCmd.Cmd
		subErr := ErrPipelineRspEmpty
		if idx > 0 {
			// pipeline命令只上报一次
			pipelienStatVal = PipeLineCmdStatValZero
		}

		if rspLen > 0 && idx < rspLen {
			subErr = response.CmdReply[idx].Err
		}
		pipelineReport(i.Target, i.CallerServerName, nowTime, err, cmdName, subErr, pipelienStatVal)
	}
}

func (i *impDao) getService() {
	// redis服务名,可以是options中指定的target或者配置文件/options配置的serviceName
	if i.Target == "" {
		_ = i.Option.LoadClientConfig(i.ServiceName)
		i.Target = i.Option.Target
	}

	// 主调服务名,没有则使用传递的serviceName
	if i.CallerServerName == "" {
		i.CallerServerName = i.ServiceName

		cfg := trpc.GlobalConfig()
		if cfg.Server.App != "" && cfg.Server.Server != "" {
			i.CallerServerName = cfg.Server.App + "." + cfg.Server.Server
		}
	}
}
