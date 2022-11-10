package redisx

import (
	"context"

	redigo "github.com/gomodule/redigo/redis"
)

const (
	RedisHash   = 1
	RedisZset   = 2
	RedisString = 3
	RedisList   = 4
	RedisSet    = 5
	RedisKey    = 6
)

var (
	commandMap = map[string]int{
		HashHset:         1,
		HashHsetnx:       1,
		HashHget:         1,
		HashHexists:      1,
		HashHdel:         1,
		HashHlen:         1,
		HashHstrlen:      1,
		HashHincrby:      1,
		HashHincrbyfloat: 1,
		HashHmset:        1,
		HashHmget:        1,
		HashHkeys:        1,
		HashHvals:        1,
		HashHgetall:      1,

		ZsetZadd:             2,
		ZsetZcard:            2,
		ZsetZcount:           2,
		ZsetZlexcount:        2,
		ZsetZincrby:          2,
		ZsetZmscore:          2,
		ZsetZpopmax:          2,
		ZsetZpopmin:          2,
		ZsetZrange:           2,
		ZsetZrangebyscore:    2,
		ZsetZrank:            2,
		ZsetZrem:             2,
		ZsetZremrangebyrank:  2,
		ZsetZremrangebyscore: 2,
		ZsetZrevrange:        2,
		ZsetZrevrangebyscore: 2,
		ZsetZrevrank:         2,
		ZsetZscore:           2,
		ZsetZrandmember:      2,

		StringSet:         3,
		StringSetex:       3,
		StringPsetex:      3,
		StringSetrange:    3,
		StringGet:         3,
		StringGetrange:    3,
		StringGetset:      3,
		StringIncr:        3,
		StringIncrby:      3,
		StringIncrbyfloat: 3,
		StringDecr:        3,
		StringDecrby:      3,
		StringStrlen:      3,
		StringAppend:      3,
		StringMget:        3,
		StringMset:        3,
		StringMsetnx:      3,

		ListLpush:      4,
		ListLpushx:     4,
		ListLrange:     4,
		ListLrem:       4,
		ListLset:       4,
		ListLtrim:      4,
		ListRpop:       4,
		ListRpoplpush:  4,
		ListRpush:      4,
		ListRpushx:     4,
		ListBlpop:      4,
		ListBrpop:      4,
		ListBrpoplpush: 4,
		ListLindex:     4,
		ListLinsert:    4,
		ListLlen:       4,
		ListLpop:       4,

		SetSadd:        5,
		SetScard:       5,
		SetSdiff:       5,
		SetSdiffstore:  5,
		SetSinter:      5,
		SetSinterstore: 5,
		SetSismember:   5,
		SetSmembers:    5,
		SetSmove:       5,
		SetSpop:        5,
		SetSrandmember: 5,
		SetSrem:        5,
		SetSunion:      5,
		SetSunionstore: 5,

		KeyDel:       6,
		KeyExists:    6,
		KeyExpire:    6,
		KeyExpireat:  6,
		KeyKeys:      6,
		KeyPersist:   6,
		KeyPexpire:   6,
		KeyPexpireat: 6,
		KeyPttl:      6,
		KeyRandomkey: 6,
		KeyTouch:     6,
		KeyTTL:       6,
		KeyType:      6,
	}
)

const (
	OK   = "OK"
	Test = "test"
)

type MyProxy struct {
	Client
}

func (MyProxy) GetRedisType(cmd string) int64 {
	typ, ok := commandMap[cmd]
	if ok {
		return int64(typ)
	}

	return 0
}

func (MyProxy) Hash(ctx context.Context,
	cmd string, args ...interface{}) (interface{}, error) {
	switch cmd {
	case HashHset:
		return int64(7), nil
	case HashHsetnx:
		return int64(1), nil
	case HashHget:
		return "111", nil
	case HashHexists:
		return int64(111), nil
	case HashHdel:
		return int64(8), nil
	case HashHlen:
		return int64(9), nil
	case HashHstrlen:
		return int64(10), nil
	case HashHincrby:
		return int64(11), nil
	case HashHincrbyfloat:
		var data []byte = []byte("2.2")
		return data, nil
	case HashHmset:
		return OK, nil
	case HashHmget:
		return []interface{}{"1", "2", "3", "4"}, nil
	case HashHkeys:
		return []interface{}{"1", "2", "3", "4", "5"}, nil
	case HashHvals:
		return []interface{}{"1", "2", "3", "4", "5"}, nil
	case HashHgetall:
		var data1 []byte = []byte("1")
		var data2 []byte = []byte("2")
		var data3 []byte = []byte("3")
		var data4 []byte = []byte("4")
		return []interface{}{data1, data2, data3, data4}, nil
	default:
		return nil, nil
	}
}

func (MyProxy) Zset(ctx context.Context,
	cmd string, args ...interface{}) (interface{}, error) {
	switch cmd {
	case ZsetZadd:
		if len(args) >= 2 {
			if args[1] == "incr" {
				var data []byte = []byte("2.2")
				return data, nil
			}
		}
		return int64(12), nil
	case ZsetZcard, ZsetZcount, ZsetZlexcount:
		return int64(15), nil
	case ZsetZincrby:
		var data []byte = []byte("2.2")
		return data, nil
	case ZsetZmscore:
		var data1 []byte = []byte("1")
		var data2 []byte = []byte("2")
		var data3 []byte = []byte("3")
		var data4 []byte = []byte("4")
		return []interface{}{data1, data2, data3, data4}, nil
	case ZsetZpopmax:
		var data1 string
		var data2 []byte = []byte("2")
		return []interface{}{data1, data2}, nil
	case ZsetZpopmin:
		var data1 string
		var data2 []byte = []byte("2")
		return []interface{}{data1, data2}, nil
	case ZsetZrange:
		var data []byte = []byte("2")
		return []interface{}{"1", data}, nil
	case ZsetZrangebyscore:
		var data []byte = []byte("2")
		return []interface{}{"1", data}, nil
	case ZsetZrank, ZsetZrem:
		return int64(16), nil
	case ZsetZremrangebyrank:
		return int64(18), nil
	case ZsetZremrangebyscore:
		return int64(19), nil
	case ZsetZrevrange:
		var data []byte = []byte("2")
		return []interface{}{"1", data}, nil
	case ZsetZrevrangebyscore:
		var data []byte = []byte("2")
		return []interface{}{"1", data}, nil
	case ZsetZrevrank:
		return int64(16), nil
	case ZsetZscore:
		var data []byte = []byte("2.2")
		return data, nil
	case ZsetZrandmember:
		return []interface{}{"1", "2"}, nil
	default:
		return nil, nil
	}
}

func (MyProxy) String(ctx context.Context,
	cmd string, args ...interface{}) (interface{}, error) {
	switch cmd {
	case StringSet:
		return OK, nil
	case StringSetex:
		return OK, nil
	case StringPsetex:
		return OK, nil
	case StringSetrange:
		return int64(11), nil
	case StringGet:
		return Test, nil
	case StringGetrange:
		return Test, nil
	case StringGetset:
		return Test, nil
	case StringIncr:
		return int64(1), nil
	case StringIncrby:
		return int64(2), nil
	case StringIncrbyfloat:
		var data []byte = []byte("2.2")
		return data, nil
	case StringDecr:
		return int64(3), nil
	case StringDecrby:
		return int64(4), nil
	case StringStrlen:
		return int64(5), nil
	case StringAppend:
		return int64(6), nil
	case StringMget:
		return []interface{}{"11", "22", "33"}, nil
	case StringMset:
		return OK, nil
	case StringMsetnx:
		return OK, nil
	default:
		return nil, nil
	}
}

func (MyProxy) List(ctx context.Context,
	cmd string, args ...interface{}) (interface{}, error) {
	switch cmd {
	case ListLpush:
		return int64(20), nil
	case ListLpushx:
		return int64(21), nil
	case ListLrange:
		return []interface{}{"1", "2"}, nil
	case ListLrem:
		return int64(22), nil
	case ListLset:
		return OK, nil
	case ListLtrim:
		return OK, nil
	case ListRpop:
		return "1", nil
	case ListRpoplpush:
		return "2", nil
	case ListRpush:
		return int64(23), nil
	case ListRpushx:
		return int64(24), nil
	case ListBlpop:
		return []interface{}{"1", "2"}, nil
	case ListBrpop:
		return []interface{}{"1", "2"}, nil
	case ListBrpoplpush:
		return "1", nil
	case ListLindex:
		return "1", nil
	case ListLinsert:
		return int64(25), nil
	case ListLlen:
		return int64(26), nil
	case ListLpop:
		return "3", nil
	default:
		return nil, nil
	}
}

func (MyProxy) Set(ctx context.Context,
	cmd string, args ...interface{}) (interface{}, error) {
	switch cmd {
	case SetSadd:
		return int64(27), nil
	case SetScard:
		return int64(28), nil
	case SetSdiff:
		return []interface{}{"11", "22", "33"}, nil
	case SetSdiffstore:
		return int64(29), nil
	case SetSinter:
		return []interface{}{"11", "22", "33"}, nil
	case SetSinterstore:
		return int64(30), nil
	case SetSismember:
		return int64(1), nil
	case SetSmembers:
		return []interface{}{"11", "22", "33"}, nil
	case SetSmove:
		return int64(1), nil
	case SetSpop:
		return "1", nil
	case SetSrandmember:
		return "1", nil
	case SetSrem:
		return int64(31), nil
	case SetSunion:
		return []interface{}{"11", "22", "33"}, nil
	case SetSunionstore:
		return int64(31), nil
	default:
		return nil, nil
	}
}

func (MyProxy) Key(ctx context.Context,
	cmd string, args ...interface{}) (interface{}, error) {
	switch cmd {
	case KeyDel:
		return int64(32), nil
	case KeyExists:
		return int64(1), nil
	case KeyExpire:
		return int64(1), nil
	case KeyExpireat:
		return int64(1), nil
	case KeyKeys:
		return []interface{}{"1", "2"}, nil
	case KeyPersist:
		return int64(1), nil
	case KeyPexpire:
		return int64(1), nil
	case KeyPexpireat:
		return int64(1), nil
	case KeyPttl:
		return int64(0), nil
	case KeyRandomkey:
		return "1111", nil
	case KeyTouch:
		return int64(33), nil
	case KeyTTL:
		return int64(0), nil
	case KeyType:
		return "set", nil
	default:
		return nil, nil
	}
}

func (mp MyProxy) Do(ctx context.Context,
	cmd string, args ...interface{}) (interface{}, error) {

	typ := mp.GetRedisType(cmd)
	switch typ {
	case RedisHash:
		return mp.Hash(ctx, cmd, args...)
	case RedisZset:
		return mp.Zset(ctx, cmd, args...)
	case RedisString:
		return mp.String(ctx, cmd, args...)
	case RedisList:
		return mp.List(ctx, cmd, args...)
	case RedisSet:
		return mp.Set(ctx, cmd, args...)
	case RedisKey:
		return mp.Key(ctx, cmd, args...)
	default:
		return nil, nil
	}

}

func (MyProxy) Pipeline(ctx context.Context) (redigo.Conn, error) {
	return nil, nil
}
