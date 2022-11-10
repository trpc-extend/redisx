package redisx

import (
	"context"

	"git.code.oa.com/trpc-go/trpc-database/redis"
)

// SADD key member [member ...]
func (i *impDao) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(members...)

	reply, err := i.Do(ctx, SetSadd, args...)
	return Int64Reply(reply, err)
}

// SCARD key
func (i *impDao) SCard(ctx context.Context, key string) (int64, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, SetScard, args...)
	return Int64Reply(reply, err)
}

// SDIFF key [key ...]
func (i *impDao) SDiff(ctx context.Context, keys ...string) ([]string, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, SetSdiff, args...)
	return StringSliceReply(reply, err)
}

// SDIFFSTORE destination key [key ...]
func (i *impDao) SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	args := redis.Args{}
	args = args.Add(destination)
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, SetSdiffstore, args...)
	return Int64Reply(reply, err)
}

// SINTER key [key ...]
func (i *impDao) SInter(ctx context.Context, keys ...string) ([]string, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, SetSinter, args...)
	return StringSliceReply(reply, err)
}

// SINTERSTORE destination key [key ...]
func (i *impDao) SInterStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	args := redis.Args{}
	args = args.Add(destination)
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, SetSinterstore, args...)
	return Int64Reply(reply, err)
}

// SISMEMBER key member
func (i *impDao) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	args := NewCmdArgs(key, member)

	reply, err := i.Do(ctx, SetSismember, args...)
	return BoolReply(reply, err)
}

// SMEMBERS key
func (i *impDao) SMembers(ctx context.Context, key string) ([]string, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, SetSmembers, args...)
	return StringSliceReply(reply, err)
}

// SMOVE source destination member
func (i *impDao) SMove(ctx context.Context, source, destination string, member interface{}) (bool, error) {
	args := NewCmdArgs(source, destination, member)

	reply, err := i.Do(ctx, SetSmove, args...)
	return BoolReply(reply, err)
}

// SPOP key
func (i *impDao) SPop(ctx context.Context, key string) (string, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, SetSpop, args...)
	return StringReply(reply, err)
}

// SPOP key count
func (i *impDao) SPopN(ctx context.Context, key string, count int64) ([]string, error) {
	args := NewCmdArgs(key, count)

	reply, err := i.Do(ctx, SetSpop, args...)
	return StringSliceReply(reply, err)
}

// SRANDMEMBER key
func (i *impDao) SRandMember(ctx context.Context, key string) (string, error) {
	args := NewCmdArgs(key)

	reply, err := i.Do(ctx, SetSrandmember, args...)
	return StringReply(reply, err)
}

// SRANDMEMBER key count
func (i *impDao) SRandMemberN(ctx context.Context, key string, count int64) ([]string, error) {
	args := NewCmdArgs(key, count)

	reply, err := i.Do(ctx, SetSrandmember, args...)
	return StringSliceReply(reply, err)
}

// SREM key member [member ...]
func (i *impDao) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(members...)

	reply, err := i.Do(ctx, SetSrem, args...)
	return Int64Reply(reply, err)
}

// SUNION key [key ...]
func (i *impDao) SUnion(ctx context.Context, keys ...string) ([]string, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, SetSunion, args...)
	return StringSliceReply(reply, err)
}

// SUNIONSTORE destination key [key ...]
func (i *impDao) SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	args := redis.Args{}
	args = args.Add(destination)
	for _, key := range keys {
		args = args.Add(key)
	}

	reply, err := i.Do(ctx, SetSunionstore, args...)
	return Int64Reply(reply, err)
}
