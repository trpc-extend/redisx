package redisx

import (
	"context"
	"time"

	"git.code.oa.com/trpc-go/trpc-database/redis"
)

// DEL key [key ...]
func (i *impDao) Del(ctx context.Context, keys ...string) (int64, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, KeyDel, args...)
	return Int64Reply(reply, err)
}

// EXISTS key [key ...]
func (i *impDao) Exists(ctx context.Context, keys ...string) (int64, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, KeyExists, args...)
	return Int64Reply(reply, err)
}

// EXPIRE key seconds
func (i *impDao) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	args := NewCmdArgs(key, FormatSec(expiration))

	reply, err := i.Do(ctx, KeyExpire, args...)
	return BoolReply(reply, err)
}

// EXPIREAT key timestamp
func (i *impDao) ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	args := NewCmdArgs(key, tm.Unix())

	reply, err := i.Do(ctx, KeyExpireat, args...)
	return BoolReply(reply, err)
}

// KEYS pattern
func (i *impDao) Keys(ctx context.Context, pattern string) ([]string, error) {
	args := NewCmdArgs(pattern)

	reply, err := i.Do(ctx, KeyKeys, args...)
	return StringSliceReply(reply, err)
}

// PERSIST key
func (i *impDao) Persist(ctx context.Context, key string) (bool, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, KeyPersist, args...)
	return BoolReply(reply, err)
}

// PEXPIRE key milliseconds
func (i *impDao) PExpire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	args := NewCmdArgs(key, FormatMs(expiration))

	reply, err := i.Do(ctx, KeyPexpire, args...)
	return BoolReply(reply, err)
}

// PEXPIREAT key milliseconds-timestamp
func (i *impDao) PExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	args := NewCmdArgs(key, tm.UnixNano()/int64(time.Millisecond))

	reply, err := i.Do(ctx, KeyPexpireat, args...)
	return BoolReply(reply, err)
}

// PTTL key
func (i *impDao) PTTL(ctx context.Context, key string) (time.Duration, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, KeyPttl, args...)
	return DurationReply(time.Millisecond, reply, err)
}

// RANDOMKEY
func (i *impDao) RandomKey(ctx context.Context) (string, error) {
	reply, err := i.Do(ctx, KeyRandomkey)
	return StringReply(reply, err)
}

// TOUCH key [key ...]
func (i *impDao) Touch(ctx context.Context, keys ...string) (int64, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, KeyTouch, args...)
	return Int64Reply(reply, err)
}

// TTL key
func (i *impDao) TTL(ctx context.Context, key string) (time.Duration, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, KeyTTL, args...)
	return DurationReply(time.Second, reply, err)
}

// TYPE key
func (i *impDao) Type(ctx context.Context, key string) (string, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, KeyType, args...)
	return StringReply(reply, err)
}
