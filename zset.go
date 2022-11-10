package redisx

import (
	"context"
)

// Z represents sorted set member.
type Z struct {
	Score  float64
	Member interface{}
}

// ZRangeBy struct
type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

// zset flag
const (
	ZsetFlagXX         = "xx"
	ZsetFlagNX         = "nx"
	ZsetFlagCH         = "ch"
	ZsetFlagIncr       = "incr"
	ZsetFlagWithscores = "withscores"
	ZsetFlagLimit      = "limit"
)

// zAdd function
func (i *impDao) zAdd(ctx context.Context, args []interface{}, n int, members ...*Z) (int64, error) {
	for i, m := range members {
		args[n+2*i] = m.Score
		args[n+2*i+1] = m.Member
	}

	reply, err := i.Do(ctx, ZsetZadd, args...)
	return Int64Reply(reply, err)
}

// ZADD key score member [score member ...]
func (i *impDao) ZAdd(ctx context.Context, key string, members ...*Z) (int64, error) {
	const n = 1
	args := make([]interface{}, n+len(members)*2)
	args[0] = key

	return i.zAdd(ctx, args, n, members...)
}

// ZADD key NX score member [score member ...]
func (i *impDao) ZAddNX(ctx context.Context, key string, members ...*Z) (int64, error) {
	const n = 2
	args := make([]interface{}, n+len(members)*2)
	args[0] = key
	args[1] = ZsetFlagNX

	return i.zAdd(ctx, args, n, members...)
}

// ZADD key XX score member [score member ...]
func (i *impDao) ZAddXX(ctx context.Context, key string, members ...*Z) (int64, error) {
	const n = 2
	args := make([]interface{}, n+len(members)*2)
	args[0] = key
	args[1] = ZsetFlagXX

	return i.zAdd(ctx, args, n, members...)
}

// ZADD key CH score member [score member ...]
func (i *impDao) ZAddCh(ctx context.Context, key string, members ...*Z) (int64, error) {
	const n = 2
	args := make([]interface{}, n+len(members)*2)
	args[0] = key
	args[1] = ZsetFlagCH

	return i.zAdd(ctx, args, n, members...)
}

// ZADD key NX CH score member [score member ...]
func (i *impDao) ZAddNXCh(ctx context.Context, key string, members ...*Z) (int64, error) {
	const n = 3
	args := make([]interface{}, n+len(members)*2)
	args[0] = key
	args[1] = ZsetFlagNX
	args[2] = ZsetFlagCH

	return i.zAdd(ctx, args, n, members...)
}

// ZADD key XX CH score member [score member ...]
func (i *impDao) ZAddXXCh(ctx context.Context, key string, members ...*Z) (int64, error) {
	const n = 3
	args := make([]interface{}, n+len(members)*2)
	args[0] = key
	args[1] = ZsetFlagXX
	args[2] = ZsetFlagCH

	return i.zAdd(ctx, args, n, members...)
}

// ZADD key INCR score member
func (i *impDao) ZIncr(ctx context.Context, key string, member Z) (float64, error) {
	args := []interface{}{key, ZsetFlagIncr, member.Score, member.Member}

	reply, err := i.Do(ctx, ZsetZadd, args...)
	return FloatReply(reply, err)
}

// ZADD key NX INCR score member
func (i *impDao) ZIncrNX(ctx context.Context, key string, member Z) (float64, error) {
	args := []interface{}{key, ZsetFlagIncr, ZsetFlagNX, member.Score, member.Member}

	reply, err := i.Do(ctx, ZsetZadd, args...)
	return FloatReply(reply, err)
}

// ZADD key XX INCR score member
func (i *impDao) ZIncrXX(ctx context.Context, key string, member Z) (float64, error) {
	args := []interface{}{key, ZsetFlagIncr, ZsetFlagXX, member.Score, member.Member}

	reply, err := i.Do(ctx, ZsetZadd, args...)
	return FloatReply(reply, err)
}

// ZCARD key
func (i *impDao) ZCard(ctx context.Context, key string) (int64, error) {
	reply, err := i.Do(ctx, ZsetZcard, key)
	return Int64Reply(reply, err)
}

// ZCOUNT key min max
func (i *impDao) ZCount(ctx context.Context, key, min, max string) (int64, error) {
	reply, err := i.Do(ctx, ZsetZcount, key, min, max)
	return Int64Reply(reply, err)
}

// ZLEXCOUNT key min max
func (i *impDao) ZLexCount(ctx context.Context, key, min, max string) (int64, error) {
	reply, err := i.Do(ctx, ZsetZlexcount, key, min, max)
	return Int64Reply(reply, err)
}

// ZINCRBY key increment member
func (i *impDao) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	reply, err := i.Do(ctx, ZsetZincrby, key, increment, member)
	return FloatReply(reply, err)
}

// ZMSCORE key member [member ...]
func (i *impDao) ZMScore(ctx context.Context, key string, members ...string) ([]interface{}, error) {
	args := make([]interface{}, 1+len(members))
	args[0] = key
	for i, m := range members {
		args[1+i] = m
	}

	reply, err := i.Do(ctx, ZsetZmscore, args...)
	return SliceReply(reply, err)
}

// ZPOPMAX key [count]
func (i *impDao) ZPopMax(ctx context.Context, key string, count ...int64) ([]Z, error) {
	args := []interface{}{
		key,
	}
	switch len(count) {
	case 0:
		break
	case 1:
		args = append(args, count[0])
	default:
		panic("too many arguments")
	}

	reply, err := i.Do(ctx, ZsetZpopmax, args...)
	if err != nil {
		return nil, err
	}

	return ZSliceReply(reply, err)
}

// ZPOPMIN key [count]
func (i *impDao) ZPopMin(ctx context.Context, key string, count ...int64) ([]Z, error) {
	args := []interface{}{
		key,
	}
	switch len(count) {
	case 0:
		break
	case 1:
		args = append(args, count[0])
	default:
		panic("too many arguments")
	}

	reply, err := i.Do(ctx, ZsetZpopmin, args...)
	if err != nil {
		return nil, err
	}

	return ZSliceReply(reply, err)
}

// ZRANGE key min max
func (i *impDao) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	reply, err := i.Do(ctx, ZsetZrange, key, start, stop)
	return StringSliceReply(reply, err)
}

// ZRANGE key min max WITHSCORES
func (i *impDao) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	reply, err := i.Do(ctx, ZsetZrange, key, start, stop, ZsetFlagWithscores)
	return ZSliceReply(reply, err)
}

// ZRANGEBYSCORE key min max [LIMIT offset count]
func (i *impDao) ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) ([]string, error) {
	args := []interface{}{key, opt.Min, opt.Max}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, ZsetFlagLimit, opt.Offset, opt.Count)
	}
	reply, err := i.Do(ctx, ZsetZrangebyscore, args...)
	return StringSliceReply(reply, err)
}

// ZRANGEBYSCORE key min max WITHSCORES [LIMIT offset count]
func (i *impDao) ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) ([]Z, error) {
	args := []interface{}{key, opt.Min, opt.Max, ZsetFlagWithscores}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, ZsetFlagLimit, opt.Offset, opt.Count)
	}
	reply, err := i.Do(ctx, ZsetZrangebyscore, args...)
	return ZSliceReply(reply, err)
}

// ZRANK key member
func (i *impDao) ZRank(ctx context.Context, key, member string) (int64, error) {
	reply, err := i.Do(ctx, ZsetZrank, key, member)
	return Int64Reply(reply, err)
}

// ZREM key member [member ...]
func (i *impDao) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	args := make([]interface{}, 1+len(members))
	args[0] = key
	for i, m := range members {
		args[1+i] = m
	}
	reply, err := i.Do(ctx, ZsetZrem, args...)
	return Int64Reply(reply, err)
}

// ZREMRANGEBYRANK key start stop
func (i *impDao) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	reply, err := i.Do(ctx, ZsetZremrangebyrank, key, start, stop)
	return Int64Reply(reply, err)
}

// ZREMRANGEBYSCORE key min max
func (i *impDao) ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error) {
	reply, err := i.Do(ctx, ZsetZremrangebyscore, key, min, max)
	return Int64Reply(reply, err)
}

// ZREVRANGE key start stop
func (i *impDao) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	reply, err := i.Do(ctx, ZsetZrevrange, key, start, stop)
	return StringSliceReply(reply, err)
}

// ZREVRANGE key start stop WITHSCORES
func (i *impDao) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	reply, err := i.Do(ctx, ZsetZrevrange, key, start, stop, ZsetFlagWithscores)
	return ZSliceReply(reply, err)
}

// ZREVRANGEBYSCORE key max min [LIMIT offset count]
func (i *impDao) ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) ([]string, error) {
	args := []interface{}{key, opt.Max, opt.Min}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, ZsetFlagLimit, opt.Offset, opt.Count)
	}
	reply, err := i.Do(ctx, ZsetZrevrangebyscore, args...)
	return StringSliceReply(reply, err)
}

// ZREVRANGEBYSCORE key max min WITHSCORES [LIMIT offset count]
func (i *impDao) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) ([]Z, error) {
	args := []interface{}{key, opt.Max, opt.Min, ZsetFlagWithscores}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(args, ZsetFlagLimit, opt.Offset, opt.Count)
	}
	reply, err := i.Do(ctx, ZsetZrevrangebyscore, args...)
	return ZSliceReply(reply, err)
}

// ZREVRANK key member
func (i *impDao) ZRevRank(ctx context.Context, key, member string) (int64, error) {
	reply, err := i.Do(ctx, ZsetZrevrank, key, member)
	return Int64Reply(reply, err)
}

// ZSCORE key member
func (i *impDao) ZScore(ctx context.Context, key, member string) (float64, error) {
	reply, err := i.Do(ctx, ZsetZscore, key, member)
	return FloatReply(reply, err)
}

// ZRANDMEMBER key count [WITHSCORES]
func (i *impDao) ZRandMember(ctx context.Context, key string, count int, withScores bool) ([]string, error) {
	args := []interface{}{key, count}
	if withScores {
		args = append(args, ZsetFlagWithscores)
	}

	reply, err := i.Do(ctx, ZsetZrandmember, args)
	return StringSliceReply(reply, err)
}
