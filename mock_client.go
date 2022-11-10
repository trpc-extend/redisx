package redisx

import (
	"context"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

// MockClient mock client
type MockClient struct {
}

// Do mock
func (*MockClient) Do(ctx context.Context, cmd string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

// RetryBackoff mock
func (*MockClient) RetryBackoff(retry int, minBackoff, maxBackoff time.Duration) {

}

// Pipeline mock
func (*MockClient) Pipeline(ctx context.Context, request *PipelineRequest) (*PipelineResponse, error) {
	return nil, nil
}

// GetPipelineConn mock
func (*MockClient) GetPipelineConn(ctx context.Context) (redigo.Conn, error) {
	var conn redigo.Conn
	return conn, nil
}

// Script mock
func (*MockClient) Script(ctx context.Context,
	keyCount int, script string, keysAndArgs ...interface{}) (interface{}, error) {
	return nil, nil
}

// Set mock
func (*MockClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

// SetNX mock
func (*MockClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return true, nil
}

// SetEX mock
func (*MockClient) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

// PSetEX mock
func (*MockClient) PSetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return nil
}

// SetXX mock
func (*MockClient) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return true, nil
}

// SetRange mock
func (*MockClient) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	return 0, nil
}

// Get mock
func (*MockClient) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

// GetRange mock
func (*MockClient) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	return "", nil
}

// GetSet mock
func (*MockClient) GetSet(ctx context.Context, key string, value interface{}) (string, error) {
	return "", nil
}

// Incr mock
func (*MockClient) Incr(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

// IncrBy mock
func (*MockClient) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return 0, nil
}

// IncrByFloat mock
func (*MockClient) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	return 0.0, nil
}

// Decr mock
func (*MockClient) Decr(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

// DecrBy mock
func (*MockClient) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	return 0, nil
}

// StrLen mock
func (*MockClient) StrLen(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

// Append mock
func (*MockClient) Append(ctx context.Context, key, value string) (int64, error) {
	return 0, nil
}

// MGet mock
func (*MockClient) MGet(ctx context.Context, keys ...string) ([]string, error) {
	return []string{}, nil
}

// MSet mock
func (*MockClient) MSet(ctx context.Context, values ...interface{}) error {
	return nil
}

// MSetNX mock
func (*MockClient) MSetNX(ctx context.Context, values ...interface{}) (bool, error) {
	return true, nil
}

// BLPop mock
func (*MockClient) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return []string{}, nil
}

// BRPop mock
func (*MockClient) BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return []string{}, nil
}

// BRPopLPush mock
func (*MockClient) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error) {
	return "", nil
}

// LIndex mock
func (*MockClient) LIndex(ctx context.Context, key string, index int64) (string, error) {
	return "", nil
}

// LInsert mock
func (*MockClient) LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error) {
	return 0, nil
}

// LInsertBefore mock
func (*MockClient) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return 0, nil
}

// LInsertAfter mock
func (*MockClient) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return 0, nil
}

// LLen mock
func (*MockClient) LLen(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

// LPop mock
func (*MockClient) LPop(ctx context.Context, key string) (string, error) {
	return "", nil
}

// LPush mock
func (*MockClient) LPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return 0, nil
}

// LPushX mock
func (*MockClient) LPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return 0, nil
}

// LRange mock
func (*MockClient) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return []string{}, nil
}

// LRem mock
func (*MockClient) LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	return 0, nil
}

// LSet mock
func (*MockClient) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	return nil
}

// LTrim mock
func (*MockClient) LTrim(ctx context.Context, key string, start, stop int64) error {
	return nil
}

// RPop mock
func (*MockClient) RPop(ctx context.Context, key string) (string, error) {
	return "", nil
}

// RPopLPush mock
func (*MockClient) RPopLPush(ctx context.Context, source, destination string) (string, error) {
	return "", nil
}

// RPush mock
func (*MockClient) RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return 0, nil
}

// RPushX mock
func (*MockClient) RPushX(ctx context.Context, key string, value interface{}) (int64, error) {
	return 0, nil
}

// SAdd mock
func (*MockClient) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return 0, nil
}

// SCard mock
func (*MockClient) SCard(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

// SDiff mock
func (*MockClient) SDiff(ctx context.Context, keys ...string) ([]string, error) {
	return []string{}, nil
}

// SDiffStore mock
func (*MockClient) SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return 0, nil
}

// SInter mock
func (*MockClient) SInter(ctx context.Context, keys ...string) ([]string, error) {
	return []string{}, nil
}

// SInterStore mock
func (*MockClient) SInterStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return 0, nil
}

// SIsMember mock
func (*MockClient) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return true, nil
}

// SMembers mock
func (*MockClient) SMembers(ctx context.Context, key string) ([]string, error) {
	return []string{}, nil
}

// SMove mock
func (*MockClient) SMove(ctx context.Context, source, destination string, member interface{}) (bool, error) {
	return true, nil
}

// SPop mock
func (*MockClient) SPop(ctx context.Context, key string) (string, error) {
	return "", nil
}

// SPopN mock
func (*MockClient) SPopN(ctx context.Context, key string, count int64) ([]string, error) {
	return []string{}, nil
}

// SRandMember mock
func (*MockClient) SRandMember(ctx context.Context, key string) (string, error) {
	return "", nil
}

// SRandMemberN mock
func (*MockClient) SRandMemberN(ctx context.Context, key string, count int64) ([]string, error) {
	return []string{}, nil
}

// SRem mock
func (*MockClient) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return 0, nil
}

// SUnion mock
func (*MockClient) SUnion(ctx context.Context, keys ...string) ([]string, error) {
	return []string{}, nil
}

// SUnionStore mock
func (*MockClient) SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return 0, nil
}

// Del mock
func (*MockClient) Del(ctx context.Context, keys ...string) (int64, error) {
	return 0, nil
}

// Exists mock
func (*MockClient) Exists(ctx context.Context, keys ...string) (int64, error) {
	return 0, nil
}

// Expire mock
func (*MockClient) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return true, nil
}

// ExpireAt mock
func (*MockClient) ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	return true, nil
}

// Keys mock
func (*MockClient) Keys(ctx context.Context, pattern string) ([]string, error) {
	return []string{}, nil
}

// Persist mock
func (*MockClient) Persist(ctx context.Context, key string) (bool, error) {
	return true, nil
}

// PExpire mock
func (*MockClient) PExpire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return true, nil
}

// PExpireAt mock
func (*MockClient) PExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	return true, nil
}

// PTTL mock
func (*MockClient) PTTL(ctx context.Context, key string) (time.Duration, error) {
	var t time.Duration
	return t, nil
}

// RandomKey mock
func (*MockClient) RandomKey(ctx context.Context) (string, error) {
	return "", nil
}

// Touch mock
func (*MockClient) Touch(ctx context.Context, keys ...string) (int64, error) {
	return 0, nil
}

// TTL mock
func (*MockClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	var t time.Duration
	return t, nil
}

// Type mock
func (*MockClient) Type(ctx context.Context, key string) (string, error) {
	return "", nil
}

// HSet mock
func (*MockClient) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return 0, nil
}

// HSetNX mock
func (*MockClient) HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	return true, nil
}

// HGet mock
func (*MockClient) HGet(ctx context.Context, key, field string) (string, error) {
	return "", nil
}

// HExists mock
func (*MockClient) HExists(ctx context.Context, key, field string) (bool, error) {
	return true, nil
}

// HDel mock
func (*MockClient) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return 0, nil
}

// HLen mock
func (*MockClient) HLen(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

// HStrLen mock
func (*MockClient) HStrLen(ctx context.Context, key string, field string) (int64, error) {
	return 0, nil
}

// HIncrBy mock
func (*MockClient) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return 0, nil
}

// HIncrByFloat mock
func (*MockClient) HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error) {
	return 0.0, nil
}

// HMSet mock
func (*MockClient) HMSet(ctx context.Context, key string, values ...interface{}) (bool, error) {
	return true, nil
}

// HMGet mock
func (*MockClient) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return nil, nil
}

// HKeys mock
func (*MockClient) HKeys(ctx context.Context, key string) ([]string, error) {
	return []string{}, nil
}

// HVals mock
func (*MockClient) HVals(ctx context.Context, key string) ([]string, error) {
	return []string{}, nil
}

// HGetAll mock
func (*MockClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return nil, nil
}

// ZAdd mock
func (*MockClient) ZAdd(ctx context.Context, key string, members ...*Z) (int64, error) {
	return 0, nil
}

// ZAddNX mock
func (*MockClient) ZAddNX(ctx context.Context, key string, members ...*Z) (int64, error) {
	return 0, nil
}

// ZAddXX mock
func (*MockClient) ZAddXX(ctx context.Context, key string, members ...*Z) (int64, error) {
	return 0, nil
}

// ZAddCh mock
func (*MockClient) ZAddCh(ctx context.Context, key string, members ...*Z) (int64, error) {
	return 0, nil
}

// ZAddNXCh mock
func (*MockClient) ZAddNXCh(ctx context.Context, key string, members ...*Z) (int64, error) {
	return 0, nil
}

// ZAddXXCh mock
func (*MockClient) ZAddXXCh(ctx context.Context, key string, members ...*Z) (int64, error) {
	return 0, nil
}

// ZIncr mock
func (*MockClient) ZIncr(ctx context.Context, key string, member Z) (float64, error) {
	return 0, nil
}

// ZIncrNX mock
func (*MockClient) ZIncrNX(ctx context.Context, key string, member Z) (float64, error) {
	return 0, nil
}

// ZIncrXX mock
func (*MockClient) ZIncrXX(ctx context.Context, key string, member Z) (float64, error) {
	return 0, nil
}

// ZCard mock
func (*MockClient) ZCard(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

// ZCount mock
func (*MockClient) ZCount(ctx context.Context, key, min, max string) (int64, error) {
	return 0, nil
}

// ZLexCount mock
func (*MockClient) ZLexCount(ctx context.Context, key, min, max string) (int64, error) {
	return 0, nil
}

// ZIncrBy mock
func (*MockClient) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	return 0.0, nil
}

// ZMScore mock
func (*MockClient) ZMScore(ctx context.Context, key string, members ...string) ([]interface{}, error) {
	return nil, nil
}

// ZPopMax mock
func (*MockClient) ZPopMax(ctx context.Context, key string, count ...int64) ([]Z, error) {
	return nil, nil
}

// ZPopMin mock
func (*MockClient) ZPopMin(ctx context.Context, key string, count ...int64) ([]Z, error) {
	return nil, nil
}

// ZRange mock
func (*MockClient) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return nil, nil
}

// ZRangeWithScores mock
func (*MockClient) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	return nil, nil
}

// ZRangeByScore mock
func (*MockClient) ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) ([]string, error) {
	return nil, nil
}

// ZRangeByScoreWithScores mock
func (*MockClient) ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) ([]Z, error) {
	return nil, nil
}

// ZRank mock
func (*MockClient) ZRank(ctx context.Context, key, member string) (int64, error) {
	return 0, nil
}

// ZRem mock
func (*MockClient) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return 0, nil
}

// ZRemRangeByRank mock
func (*MockClient) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	return 0, nil
}

// ZRemRangeByScore mock
func (*MockClient) ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error) {
	return 0, nil
}

// ZRevRange mock
func (*MockClient) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return nil, nil
}

// ZRevRangeWithScores mock
func (*MockClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	return nil, nil
}

// ZRevRangeByScore mock
func (*MockClient) ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) ([]string, error) {
	return nil, nil
}

// ZRevRangeByScoreWithScores mock
func (*MockClient) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) ([]Z, error) {
	return nil, nil
}

// ZRevRank mock
func (*MockClient) ZRevRank(ctx context.Context, key, member string) (int64, error) {
	return 0, nil
}

// ZScore mock
func (*MockClient) ZScore(ctx context.Context, key, member string) (float64, error) {
	return 0.0, nil
}

// ZRandMember mock
func (*MockClient) ZRandMember(ctx context.Context, key string, count int, withScores bool) ([]string, error) {
	return nil, nil
}

// Lock mock
func (*MockClient) Lock(ctx context.Context, key string, expiration time.Duration) error {
	return nil
}

// UnLock mock
func (*MockClient) UnLock(ctx context.Context) {

}
