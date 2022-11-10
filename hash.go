package redisx

import (
	"context"
)

// HSET key field value [field value ...]
func (i *impDao) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	var args []interface{}
	args = append(args, key)
	args = append(args, values...)

	reply, err := i.Do(ctx, HashHset, args...)
	return Int64Reply(reply, err)
}

// HSETNX key field value
func (i *impDao) HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	reply, err := i.Do(ctx, HashHsetnx, key, field, value)
	return BoolReply(reply, err)
}

// HGET key field
func (i *impDao) HGet(ctx context.Context, key, field string) (string, error) {
	reply, err := i.Do(ctx, HashHget, key, field)
	return StringReply(reply, err)
}

// HEXISTS key field
func (i *impDao) HExists(ctx context.Context, key, field string) (bool, error) {
	reply, err := i.Do(ctx, HashHexists, key, field)
	return BoolReply(reply, err)
}

// HDEL key field [field ...]
func (i *impDao) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	var args []interface{}
	args = append(args, key)
	for _, field := range fields {
		args = append(args, field)
	}

	reply, err := i.Do(ctx, HashHdel, args...)
	return Int64Reply(reply, err)
}

// HLEN key
func (i *impDao) HLen(ctx context.Context, key string) (int64, error) {
	reply, err := i.Do(ctx, HashHlen, key)
	return Int64Reply(reply, err)
}

// HSTRLEN key field
func (i *impDao) HStrLen(ctx context.Context, key string, field string) (int64, error) {
	reply, err := i.Do(ctx, HashHstrlen, key, field)
	return Int64Reply(reply, err)
}

// HINCRBY key field increment
func (i *impDao) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	reply, err := i.Do(ctx, HashHincrby, key, field, incr)
	return Int64Reply(reply, err)
}

// HINCRBYFLOAT key field increment
func (i *impDao) HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error) {
	reply, err := i.Do(ctx, HashHincrbyfloat, key, field, incr)
	return FloatReply(reply, err)
}

// HMSET key field value [field value ...]
func (i *impDao) HMSet(ctx context.Context, key string, values ...interface{}) (bool, error) {
	var args []interface{}
	args = append(args, key)
	args = append(args, values...)

	reply, err := i.Do(ctx, HashHmset, args...)
	return BoolReply(reply, err)
}

// HMGET key field [field ...]
func (i *impDao) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	var args []interface{}
	args = append(args, key)
	for _, field := range fields {
		args = append(args, field)
	}

	reply, err := i.Do(ctx, HashHmget, args...)
	return SliceReply(reply, err)
}

// HKEYS key
func (i *impDao) HKeys(ctx context.Context, key string) ([]string, error) {
	reply, err := i.Do(ctx, HashHkeys, key)
	return StringSliceReply(reply, err)
}

// HVALS key
func (i *impDao) HVals(ctx context.Context, key string) ([]string, error) {
	reply, err := i.Do(ctx, HashHvals, key)
	return StringSliceReply(reply, err)
}

// HGETALL key
func (i *impDao) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	reply, err := i.Do(ctx, HashHgetall, key)
	return StringStringMapReply(reply, err)
}
