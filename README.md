## redisx库是什么？

基于trpc-redis对redis的5大类型(string, zset, hash, list, set)的100多个常用高频接口进行封装，便于业务直接使用，同时提供Pipeline和Script以及底层命令接口。

## redis库支持那些接口？

```go
Do(ctx context.Context, cmd string, args ...interface{}) (interface{}, error)
RetryBackoff(retry int, minBackoff, maxBackoff time.Duration)

// pipeline
Pipeline(ctx context.Context, pipeLineCmd []PipelineCmd) ([]PipelineCmdReply, error)

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
HMSet(ctx context.Context, key string, values ...interface{}) (string, error)
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
```

## 如何使用？

```go
package main

import (
    "context"
    "github.com/trpc-extend/redisx"
)

func (s *server) SayHello(ctx context.Context, req *pb.ReqBody, rsp *pb.RspBody)( err error ) {
    proxy := redisx.NewClient(ctx, "trpc.redis.WeseeCommCount.Write") // 名字写tRPC框架配置里面的名字即可
    // redis set 命令
    err := proxy.Set(ctx, "key", "value")
    if err != nil {
        log.ContextErrorf(ctx, "set failed, err = [%v]", err)
        return err
    }
    // redis mget 命令
    keys := []string{"key1", "key2", "key3"} 
    vals, err := proxy.MGet(ctx, keys...) 
    if err != nil {
        log.ContextErrorf(ctx, "mget failed, keys = [%+v], err = [%v]", keys, err)
        return err
    } 
    return nil
}
```

## 接口示例
---
```
ctx := trpc.BackgroundContext()
rdb := redisx.NewClient("trpc.redis.xxx.xxx", //service name,主要用于监控上报和配置管理，业务自行定义
    client.WithNamespace("Production"),
	client.WithTarget("redis+polaris://:xxx@xx.xxx.redis.com"),
	client.WithTimeout(time.Duration(10)*time.Millisecond))

err := rdb.Lock(ctx, "lockkey", 100)
if err != nil {
	fmt.Printf("rdb.Lock failed err=%v\r\n", err)
	return
} 
rdb.UnLock(ctx)
```


在tRPC框架组件client配置
---
```
#TRPC框架中配置
client:                                             #客户端调用的后端配置
  service:                                          #针对单个后端的配置
    - name: trpc.redis.xxx.xxx                      #redis serviceName，由业务自行定义，
      namespace: Production
      target: polaris://xxx.test.redis.com
      password: xxxxx
      timeout: 20                                   #当前这个请求最长处理时间
    - name: trpc.redis.xxx.xxx
      namespace: Production
      target: redis+polaris://:passwd@polaris_name 
      timeout: 800                                  #当前这个请求最长处理时间
  
#代码中配置    
rdb := NewClient(ctx, 
		"base_arch_redis", 
		client.WithNamespace("Production"),
		client.WithTarget("redis+polaris://:VnPnGHMLXYq@sz.weishi_base_arch.redis.com"))
```
redis proxy配置参数配置
---
```
plugins:                            #插件配置
  database:
    redis:
        max_idle: 20                # 最大空闲连接数
        max_active: 100             # 最大活跃连接数
        idle_timeout: 180000        # 连接最大空闲等待时间(单位：毫秒)
```
