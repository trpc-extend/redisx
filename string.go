package redisx

import (
	"context"
	"time"

	"git.code.oa.com/trpc-go/trpc-database/redis"
)

// string flag
const (
	StringFlagKeepttl = "keepttl"
	StringFlagPx      = "px"
	StringFlagEx      = "ex"
	StringFlagNx      = "nx"
	StringFlagXx      = "xx"
)

// SET key value [EX seconds] [PX milliseconds] [keepttl]
func (i *impDao) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	var args []interface{}

	switch expiration {
	case 0:
		args = NewCmdArgs(key, value)
	case KeepTTL:
		args = NewCmdArgs(key, value, StringFlagKeepttl)
	default:
		if UsePrecise(expiration) {
			args = NewCmdArgs(key, value, StringFlagPx, FormatMs(expiration))
		} else {
			args = NewCmdArgs(key, value, StringFlagEx, FormatSec(expiration))
		}
	}

	reply, err := i.Do(ctx, StringSet, args...)
	return StatusReply(StringSet, reply, err)
}

// SET key value [EX seconds] [PX milliseconds] [keepttl] NX
func (i *impDao) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	var args []interface{}

	switch expiration {
	case 0:
		args = NewCmdArgs(key, value, StringFlagNx)
	case KeepTTL:
		args = NewCmdArgs(key, value, StringFlagKeepttl, StringFlagNx)
	default:
		if UsePrecise(expiration) {
			args = NewCmdArgs(key, value, StringFlagPx, FormatMs(expiration), StringFlagNx)
		} else {
			args = NewCmdArgs(key, value, StringFlagEx, FormatSec(expiration), StringFlagNx)
		}
	}

	reply, err := i.Do(ctx, StringSet, args...)
	return BoolReply(reply, err)
}

// SETEX key seconds value
func (i *impDao) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	args := NewCmdArgs(key, FormatSec(expiration), value)

	reply, err := i.Do(ctx, StringSetex, args...)
	return StatusReply(StringSetex, reply, err)
}

// PSETEX key milliseconds value
func (i *impDao) PSetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	args := NewCmdArgs(key, FormatMs(expiration), value)

	reply, err := i.Do(ctx, StringPsetex, args...)
	return StatusReply(StringPsetex, reply, err)
}

// SET key value [EX seconds] [PX milliseconds] [keepttl] XX
func (i *impDao) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	var args []interface{}

	switch expiration {
	case 0:
		args = NewCmdArgs(key, value, StringFlagXx)
	case KeepTTL:
		args = NewCmdArgs(key, value, StringFlagKeepttl, StringFlagXx)
	default:
		if UsePrecise(expiration) {
			args = NewCmdArgs(key, value, StringFlagPx, FormatMs(expiration), StringFlagXx)
		} else {
			args = NewCmdArgs(key, value, StringFlagEx, FormatSec(expiration), StringFlagXx)
		}
	}

	reply, err := i.Do(ctx, StringSet, args...)
	return BoolReply(reply, err)
}

// SETRANGE key offset value
func (i *impDao) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	args := NewCmdArgs(key, offset, value)

	reply, err := i.Do(ctx, StringSetrange, args...)
	return Int64Reply(reply, err)
}

// GET key
func (i *impDao) Get(ctx context.Context, key string) (string, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, StringGet, args...)
	return StringReply(reply, err)
}

// GETRANGE key start end
func (i *impDao) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	args := NewCmdArgs(key, start, end)

	reply, err := i.Do(ctx, StringGetrange, args...)
	return StringReply(reply, err)
}

// GETSET key value
func (i *impDao) GetSet(ctx context.Context, key string, value interface{}) (string, error) {
	args := NewCmdArgs(key, value)

	reply, err := i.Do(ctx, StringGetset, args...)
	return StringReply(reply, err)
}

// INCR key
func (i *impDao) Incr(ctx context.Context, key string) (int64, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, StringIncr, args...)
	return Int64Reply(reply, err)
}

// INCRBY key increment
func (i *impDao) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	args := NewCmdArgs(key, value)

	reply, err := i.Do(ctx, StringIncrby, args...)
	return Int64Reply(reply, err)
}

// INCRBYFLOAT key increment
func (i *impDao) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	args := NewCmdArgs(key, value)

	reply, err := i.Do(ctx, StringIncrbyfloat, args...)
	return FloatReply(reply, err)
}

// DECR key
func (i *impDao) Decr(ctx context.Context, key string) (int64, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, StringDecr, args...)
	return Int64Reply(reply, err)
}

// DECRBY key decrement
func (i *impDao) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	args := NewCmdArgs(key, decrement)

	reply, err := i.Do(ctx, StringDecrby, args...)
	return Int64Reply(reply, err)
}

// STRLEN key
func (i *impDao) StrLen(ctx context.Context, key string) (int64, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, StringStrlen, args...)
	return Int64Reply(reply, err)
}

// APPEND key value
func (i *impDao) Append(ctx context.Context, key, value string) (int64, error) {
	args := NewCmdArgs(key, value)

	reply, err := i.Do(ctx, StringAppend, args...)
	return Int64Reply(reply, err)
}

// MGET key [key ...]
func (i *impDao) MGet(ctx context.Context, keys ...string) ([]string, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, StringMget, args...)
	return StringSliceReply(reply, err)
}

// MSET key value [key value …]
func (i *impDao) MSet(ctx context.Context, values ...interface{}) error {
	args := NewCmdArgs(values...)

	reply, err := i.Do(ctx, StringMset, args...)
	return StatusReply(StringMset, reply, err)
}

// MSETNX key value [key value ...]
func (i *impDao) MSetNX(ctx context.Context, values ...interface{}) (bool, error) {
	args := NewCmdArgs(values...)

	reply, err := i.Do(ctx, StringMsetnx, args...)
	return BoolReply(reply, err)
}
