package redisx

import (
	"context"
	"time"

	"git.code.oa.com/trpc-go/trpc-database/redis"
)

const (
	listFlagBefore = "before"
	listFlagAfter  = "after"
)

// LPUSH key value [value ...]
func (i *impDao) LPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(values...)

	reply, err := i.Do(ctx, ListLpush, args...)
	return Int64Reply(reply, err)

}

// LPUSHX key element [element ...]
func (i *impDao) LPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(values...)

	reply, err := i.Do(ctx, ListLpushx, args...)
	return Int64Reply(reply, err)
}

// LRANGE key start stop
func (i *impDao) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	args := NewCmdArgs(key, start, stop)

	reply, err := i.Do(ctx, ListLrange, args...)
	return StringSliceReply(reply, err)
}

// LREM key count element
func (i *impDao) LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	args := NewCmdArgs(key, count, value)

	reply, err := i.Do(ctx, ListLrem, args...)
	return Int64Reply(reply, err)
}

// LSET key index element
func (i *impDao) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	args := NewCmdArgs(key, index, value)

	reply, err := i.Do(ctx, ListLset, args...)
	return StatusReply(ListLset, reply, err)
}

// LTRIM key start stop
func (i *impDao) LTrim(ctx context.Context, key string, start, stop int64) error {
	args := NewCmdArgs(key, start, stop)

	reply, err := i.Do(ctx, ListLtrim, args...)
	return StatusReply(ListLtrim, reply, err)
}

// RPOP key
func (i *impDao) RPop(ctx context.Context, key string) (string, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, ListRpop, args...)
	return StringReply(reply, err)
}

// RPOPLPUSH source destination
func (i *impDao) RPopLPush(ctx context.Context, source, destination string) (string, error) {
	args := NewCmdArgs(source, destination)

	reply, err := i.Do(ctx, ListRpoplpush, args...)
	return StringReply(reply, err)
}

// RPUSH key element [element ...]
func (i *impDao) RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(values...)

	reply, err := i.Do(ctx, ListRpush, args...)
	return Int64Reply(reply, err)
}

// RPUSHX key element [element ...]
func (i *impDao) RPushX(ctx context.Context, key string, value interface{}) (int64, error) {
	args := NewCmdArgs(key, value)

	reply, err := i.Do(ctx, ListRpushx, args...)
	return Int64Reply(reply, err)
}

// BLPOP key [key ...] timeout
func (i *impDao) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}
	args = args.Add(FormatSec(timeout))

	reply, err := i.Do(ctx, ListBlpop, args...)
	return StringSliceReply(reply, err)
}

// BRPOP key [key ...] timeout
func (i *impDao) BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}
	args = args.Add(FormatSec(timeout))

	reply, err := i.Do(ctx, ListBrpop, args...)
	return StringSliceReply(reply, err)

}

// BRPOPLPUSH source destination timeout
func (i *impDao) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error) {
	args := NewCmdArgs(source, destination, FormatSec(timeout))

	reply, err := i.Do(ctx, ListBrpoplpush, args...)
	return StringReply(reply, err)
}

// LINDEX key index
func (i *impDao) LIndex(ctx context.Context, key string, index int64) (string, error) {
	args := NewCmdArgs(key, index)

	reply, err := i.Do(ctx, ListLindex, args...)
	return StringReply(reply, err)
}

// LINSERT key BEFORE|AFTER pivot value
func (i *impDao) LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error) {
	args := NewCmdArgs(key, op, pivot, value)

	reply, err := i.Do(ctx, ListLinsert, args...)
	return Int64Reply(reply, err)
}

// LINSERT key BEFORE pivot value
func (i *impDao) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	args := NewCmdArgs(key, listFlagBefore, pivot, value)

	reply, err := i.Do(ctx, ListLinsert, args...)
	return Int64Reply(reply, err)
}

// LINSERT key AFTER pivot value
func (i *impDao) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	args := NewCmdArgs(key, listFlagAfter, pivot, value)

	reply, err := i.Do(ctx, ListLinsert, args...)
	return Int64Reply(reply, err)
}

// LLEN key
func (i *impDao) LLen(ctx context.Context, key string) (int64, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, ListLlen, args...)
	return Int64Reply(reply, err)
}

// LPOP key
func (i *impDao) LPop(ctx context.Context, key string) (string, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, ListLpop, args...)
	return StringReply(reply, err)

}
