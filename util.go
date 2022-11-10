package redisx

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	redigo "github.com/gomodule/redigo/redis"

	"git.code.oa.com/trpc-go/trpc-go/errs"
)

// return code
const (
	ReturnSucc       = 1
	ReturnFault      = 2
	ReturnTimeout    = 3
	ReturnCachemiss  = 4
	ReturnConnfailed = 5
)

const (
	KeepTTL = -1 //redis TTL 默认值
)

var (
	RedisComponentVersion = "0.0.1" //redix基础库版本号
)

// Sleep 重试休眠
func Sleep(ctx context.Context, dur time.Duration) error {
	t := time.NewTimer(dur)
	defer t.Stop()

	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RetryBackoff 重试随机休眠间隔
func RetryBackoff(retry int, minBackoff, maxBackoff time.Duration) time.Duration {
	if retry < 0 || minBackoff == 0 {
		return 0
	}

	d := minBackoff << uint(retry)
	if d < minBackoff {
		return maxBackoff
	}

	d = minBackoff + time.Duration(rand.Int63n(int64(d)))
	if d > maxBackoff || d < minBackoff {
		d = maxBackoff
	}

	return d
}

// UsePrecise 判断Duration是否是ms
func UsePrecise(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

// FormatMs Duration 转int64类型的毫秒值
func FormatMs(dur time.Duration) int64 {
	if dur > 0 && dur < time.Millisecond {
		return 1
	}
	return int64(dur / time.Millisecond)
}

// FormatSec Duration 转int64类型的秒值
func FormatSec(dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		return 1
	}
	return int64(dur / time.Second)
}

// err2num err适配
func err2num(err error, model int) int {
	switch model {
	case ReturnSucc:
		if err == nil || err == ErrNil {
			return 1
		}
	case ReturnFault:
		if err != nil {
			return 1
		}
	case ReturnTimeout:
		if errs.Code(err) == errs.RetClientTimeout {
			return 1
		}
	case ReturnCachemiss:
		if err == ErrNil {
			return 1
		}
	case ReturnConnfailed:
		errCode := errs.Code(err)
		if errCode == errs.RetClientConnectFail || errCode == errs.RetClientNetErr {
			return 1
		}
	default:
		return 0
	}

	return 0
}

// NewCmdArgs func
func NewCmdArgs(args ...interface{}) []interface{} {
	return args
}

// replyVal 基本类型比较, 其他类型的比较暂不支持
func replyVal(val, okVal interface{}) error {
	if val == okVal {
		return nil
	}

	//TODO DeepEqual

	return errors.New("redis: commands execute failed")
}

// StatusReply parse
func StatusReply(cmd string, reply interface{}, err error) error {
	val, err := redigo.String(reply, err)
	if err == redigo.ErrNil {
		//返回值为NULL Bulk Reply,在NX,XX等条件约束下,条件未满足操作未执行
		return ErrNil
	}

	if err != nil {
		//返回异常
		return err
	}

	switch cmd {
	case StringSet, StringSetex:
		return replyVal(val, "OK")
	case "ping":
		return replyVal(val, "PONG")
	default:
		return replyVal(val, "OK")
	}
}

// BoolReply parse
func BoolReply(reply interface{}, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	switch reply := reply.(type) {
	case int64:
		return reply != 0, nil
	case []byte:
		return strconv.ParseBool(string(reply))
	case string: //这是将redis返回Simple string reply类型转换为bool类型("OK"->bool)
		return reply == "OK", nil
	case nil:
		return false, ErrNil
	default:
		return false, fmt.Errorf("redigo: unexpected type for Bool, got type %T", reply)
	}
}

// StringReply parse
func StringReply(reply interface{}, err error) (string, error) {
	val, err := redigo.String(reply, err)
	if err == redigo.ErrNil {
		return "", ErrNil
	}

	return val, err
}

// IntReply parse
func IntReply(reply interface{}, err error) (int, error) {
	return redigo.Int(reply, err)
}

// Int64Reply parse
func Int64Reply(reply interface{}, err error) (int64, error) {
	return redigo.Int64(reply, err)
}

// Uint64Reply parse
func Uint64Reply(reply interface{}, err error) (uint64, error) {
	return redigo.Uint64(reply, err)
}

// FloatReply parse
func FloatReply(reply interface{}, err error) (float64, error) {
	return redigo.Float64(reply, err)
}

// BytesReply parse
func BytesReply(reply interface{}, err error) ([]byte, error) {
	return redigo.Bytes(reply, err)
}

// SliceReply parse
func SliceReply(reply interface{}, err error) ([]interface{}, error) {
	return redigo.Values(reply, err)
}

// BoolSliceReply parse
func BoolSliceReply(reply interface{}, err error) ([]bool, error) {
	//return redigo.Bool(reply, err)
	return []bool{}, nil
}

// StringSliceReply parse
func StringSliceReply(reply interface{}, err error) ([]string, error) {
	return redigo.Strings(reply, err)
}

// IntSliceReply parse
func IntSliceReply(reply interface{}, err error) ([]int, error) {
	return redigo.Ints(reply, err)
}

// Int64SliceReply parse
func Int64SliceReply(reply interface{}, err error) ([]int64, error) {
	return redigo.Int64s(reply, err)
}

// FloatSliceReply parse
func FloatSliceReply(reply interface{}, err error) ([]float64, error) {
	return redigo.Float64s(reply, err)
}

// BytesSliceReply parse
func BytesSliceReply(reply interface{}, err error) ([][]byte, error) {
	return redigo.ByteSlices(reply, err)
}

// StringStringMapReply parse
func StringStringMapReply(reply interface{}, err error) (map[string]string, error) {
	return redigo.StringMap(reply, err)
}

// StringIntMapReply parse
func StringIntMapReply(reply interface{}, err error) (map[string]int, error) {
	return redigo.IntMap(reply, err)
}

// StringInt64MapReply parse
func StringInt64MapReply(reply interface{}, err error) (map[string]int64, error) {
	return redigo.Int64Map(reply, err)
}

// StringBoolMapReply parse
func StringBoolMapReply(reply interface{}, err error) (map[string]bool, error) {
	values, err := redigo.Values(reply, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("redigo: StringBoolMap expects even number of values result")
	}
	m := make(map[string]bool, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].([]byte)
		if !ok {
			return nil, errors.New("redigo: StringBoolMap key not a bulk string value")
		}
		value, err := redigo.Bool(values[i+1], nil)
		if err != nil {
			return nil, err
		}
		m[string(key)] = value
	}
	return m, nil
}

// DurationReply parse
func DurationReply(precision time.Duration, reply interface{}, err error) (time.Duration, error) {
	val, err := Int64Reply(reply, err)
	if err != nil {
		return time.Duration(0), err
	}

	result := time.Duration(0)
	switch val {
	// -2 if the key does not exist
	// -1 if the key exists but has no associated expire
	case -2, -1:
		result = time.Duration(val)
	default:
		result = time.Duration(val) * precision
	}
	return result, nil
}

// Sort struct
type Sort struct {
	By            string
	Offset, Count int64
	Get           []string
	Order         string
	Alpha         bool
}

func (sort *Sort) args(key string) []interface{} {
	args := []interface{}{key}
	if sort.By != "" {
		args = append(args, "by", sort.By)
	}
	if sort.Offset != 0 || sort.Count != 0 {
		args = append(args, "limit", sort.Offset, sort.Count)
	}
	for _, get := range sort.Get {
		args = append(args, "get", get)
	}
	if sort.Order != "" {
		args = append(args, sort.Order)
	}
	if sort.Alpha {
		args = append(args, "alpha")
	}
	return args
}

// ZSliceReply parse
func ZSliceReply(reply interface{}, err error) ([]Z, error) {
	lst, err := redigo.Values(reply, err)
	if err != nil {
		return nil, err
	}

	if len(lst)&1 != 0 {
		return nil, fmt.Errorf("redigo: expects even number of values result")
	}

	res := make([]Z, len(lst)/2)
	for i := 0; i < len(lst)/2; i++ {
		var tmp Z
		tmp.Member, err = StringReply(lst[i*2], nil)
		if err != nil {
			return nil, fmt.Errorf("redigo: unexpected type for String, got type %T", lst[i*2])
		}
		tmp.Score, err = FloatReply(lst[i*2+1], nil)
		if err != nil {
			return nil, fmt.Errorf("redigo: unexpected type for Float64, got type %T", lst[i*2+1])
		}
		res[i] = tmp
	}

	return res, nil
}

// ParseTarget,解析redix target获取polaris_name
func ParseTarget(name string) string {
	//target,redis+polaris://:passwd@polaris_name
	index := strings.Index(name, "@")
	if index == -1 {
		return name
	}

	target := name[index+1:]
	if target == "" {
		return name
	}

	return target
}
