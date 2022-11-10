package redisx

import (
	"context"
	"testing"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/smartystreets/goconvey/convey"

	"git.code.oa.com/NGTest/gomonkey"
	"git.code.oa.com/trpc-go/trpc-go"
)

func TestSleep(t *testing.T) {
	convey.Convey("Sleep", t, func() {
		convey.Convey("SleepNil", func() {
			ctx := trpc.BackgroundContext()
			err := Sleep(ctx, 0)
			convey.So(err, convey.ShouldBeNil)
		})
		convey.Convey("SleepNotNil", func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*100))
			defer cancel()
			err := Sleep(ctx, time.Duration(time.Millisecond*500))
			convey.So(err, convey.ShouldNotBeNil)
		})
	})
}

func TestRetryBackoff(t *testing.T) {
	convey.Convey("RetryBackoff", t, func() {
		convey.Convey("RetryBackoffBranch1", func() {
			t := RetryBackoff(-1, 0, 0)
			convey.So(t, convey.ShouldEqual, 0)
		})

		convey.Convey("RetryBackoffBranch2", func() {
			t := RetryBackoff(0, 0*time.Millisecond, 1*time.Millisecond)
			convey.So(t, convey.ShouldEqual, 0)
		})

		convey.Convey("RetryBackoffBranch3", func() {
			t := RetryBackoff(1, 1*time.Millisecond, 1*time.Millisecond)
			convey.So(t, convey.ShouldNotEqual, 0)
		})
	})
}

func TestUsePrecise(t *testing.T) {
	convey.Convey("UsePrecise", t, func() {
		convey.Convey("UsePrecise", func() {
			b := UsePrecise(0)
			convey.So(b, convey.ShouldEqual, true)
		})
	})
}

func TestFormatMs(t *testing.T) {
	convey.Convey("FormatMs", t, func() {
		convey.Convey("FormatMsBranch1", func() {
			r := FormatMs(1)
			convey.So(r, convey.ShouldEqual, 1)
		})

		convey.Convey("FormatMsBranch2", func() {
			r := FormatMs(0)
			convey.So(r, convey.ShouldNotEqual, 1)
		})
	})
}

func TestFormatSec(t *testing.T) {
	convey.Convey("FormatSec", t, func() {
		convey.Convey("FormatSecBranch1", func() {
			r := FormatSec(1)
			convey.So(r, convey.ShouldEqual, 1)
		})

		convey.Convey("FormatSecBranch2", func() {
			r := FormatSec(0)
			convey.So(r, convey.ShouldNotEqual, 1)
		})
	})
}

func TestErr2num(t *testing.T) {
	convey.Convey("err2num", t, func() {
		convey.Convey("err2numBranch1", func() {
			r := err2num(nil, 1)
			convey.So(r, convey.ShouldEqual, 1)
		})

		convey.Convey("err2numBranch2", func() {
			r := err2num(ErrNil, 2)
			convey.So(r, convey.ShouldEqual, 1)
		})

		convey.Convey("err2numBranch3", func() {
			r := err2num(nil, 3)
			convey.So(r, convey.ShouldNotEqual, 1)
		})

		convey.Convey("err2numBranch4", func() {
			r := err2num(ErrNil, 4)
			convey.So(r, convey.ShouldEqual, 1)
		})

		convey.Convey("err2numBranch5", func() {
			r := err2num(nil, 5)
			convey.So(r, convey.ShouldNotEqual, 1)
		})

		convey.Convey("err2numBranch6", func() {
			r := err2num(nil, 6)
			convey.So(r, convey.ShouldNotEqual, 1)
		})
	})
}

func TestNewCmdArgs(t *testing.T) {
	convey.Convey("NewCmdArgs", t, func() {
		r := NewCmdArgs(1)
		convey.So(len(r), convey.ShouldEqual, 1)

		r = NewCmdArgs(1, 2)
		convey.So(len(r), convey.ShouldEqual, 2)
	})
}

func TestReplyVal(t *testing.T) {
	convey.Convey("replyVal", t, func() {
		err := replyVal(1, 1)
		convey.So(err, convey.ShouldEqual, nil)

		err = replyVal(1, 2)
		convey.So(err, convey.ShouldNotEqual, nil)
	})
}

func TestStatusReply(t *testing.T) {
	convey.Convey("StatusReply", t, func() {
		convey.Convey("StatusReplyBranch1", func() {
			err := StatusReply("set", 1, nil)
			convey.So(err, convey.ShouldNotEqual, nil)

			var p1 = gomonkey.ApplyFunc(redigo.String, func(reply interface{}, err error) (string, error) {
				return "", redigo.ErrNil
			})
			defer p1.Reset()

			err = StatusReply("set", 1, nil)
			convey.So(err, convey.ShouldEqual, ErrNil)
		})

		convey.Convey("StatusReplyBranch2", func() {
			var p1 = gomonkey.ApplyFunc(redigo.String, func(reply interface{}, err error) (string, error) {
				return "", nil
			})
			defer p1.Reset()

			err := StatusReply("set", 1, nil)
			convey.So(err, convey.ShouldNotEqual, nil)

			err = StatusReply("ping", 1, nil)
			convey.So(err, convey.ShouldNotEqual, nil)

			err = StatusReply("get", 1, nil)
			convey.So(err, convey.ShouldNotEqual, nil)
		})
	})
}

func TestBoolReply(t *testing.T) {
	convey.Convey("BoolReply", t, func() {
		convey.Convey("BoolReplyBranch1", func() {
			b, err := BoolReply(nil, ErrNil)
			convey.So(b, convey.ShouldEqual, false)
			convey.So(err, convey.ShouldEqual, ErrNil)
		})
		convey.Convey("BoolReplyBranch2", func() {
			b, err := BoolReply(int64(1), nil)
			convey.So(b, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldEqual, nil)
		})
		convey.Convey("BoolReplyBranch3", func() {
			var data []byte = []byte("1")
			b, err := BoolReply(data, nil)
			convey.So(b, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldEqual, nil)
		})
		convey.Convey("BoolReplyBranch4", func() {
			b, err := BoolReply("OK", nil)
			convey.So(b, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldEqual, nil)
		})
		convey.Convey("BoolReplyBranch5", func() {
			b, err := BoolReply(0.123, nil)
			convey.So(b, convey.ShouldEqual, false)
			convey.So(err, convey.ShouldNotEqual, nil)
		})
		convey.Convey("BoolReplyBranch6", func() {
			b, err := BoolReply(nil, nil)
			convey.So(b, convey.ShouldEqual, false)
			convey.So(err, convey.ShouldNotEqual, nil)
		})
	})
}

func TestStringReply(t *testing.T) {
	convey.Convey("StringReply", t, func() {
		convey.Convey("StringReplyBranch1", func() {
			b, err := StringReply(nil, ErrNil)
			convey.So(b, convey.ShouldEqual, "")
			convey.So(err, convey.ShouldEqual, ErrNil)
		})
		convey.Convey("StringReplyBranch2", func() {
			b, err := StringReply("1", nil)
			convey.So(b, convey.ShouldEqual, "1")
			convey.So(err, convey.ShouldEqual, nil)
		})
	})
}

func TestIntReply(t *testing.T) {
	convey.Convey("IntReply", t, func() {
		b, err := IntReply(int64(123), nil)
		convey.So(b, convey.ShouldEqual, 123)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestInt64Reply(t *testing.T) {
	convey.Convey("IntReply64", t, func() {
		b, err := Int64Reply(int64(123), nil)
		convey.So(b, convey.ShouldEqual, 123)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestUint64Reply(t *testing.T) {
	convey.Convey("Uint64Reply", t, func() {
		b, err := Uint64Reply(int64(123), nil)
		convey.So(b, convey.ShouldEqual, 123)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestFloatReply(t *testing.T) {
	convey.Convey("FloatReply", t, func() {
		var data []byte = []byte("123.33")
		b, err := FloatReply(data, nil)
		convey.So(b, convey.ShouldEqual, 123.33)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestBytesReply(t *testing.T) {
	convey.Convey("BytesReply", t, func() {
		var data []byte = []byte("123.33")
		_, err := BytesReply(data, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestSliceReply(t *testing.T) {
	convey.Convey("SliceReply", t, func() {
		_, err := SliceReply(nil, nil)
		convey.So(err, convey.ShouldEqual, ErrNil)
	})
}

func TestBoolSliceReply(t *testing.T) {
	convey.Convey("BoolSliceReply", t, func() {
		_, err := BoolSliceReply(nil, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestStringSliceReply(t *testing.T) {
	convey.Convey("StringSliceReply", t, func() {
		_, err := StringSliceReply([]interface{}{"1", "2"}, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestIntSliceReply(t *testing.T) {
	convey.Convey("IntSliceReply", t, func() {
		_, err := IntSliceReply([]interface{}{int64(1), int64(2)}, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestInt64SliceReply(t *testing.T) {
	convey.Convey("Int64SliceReply", t, func() {
		_, err := Int64SliceReply([]interface{}{int64(1), int64(2)}, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestFloatSliceReply(t *testing.T) {
	convey.Convey("FloatSliceReply", t, func() {
		var data []byte = []byte("123.33")
		_, err := FloatSliceReply([]interface{}{data}, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestBytesSliceReply(t *testing.T) {
	convey.Convey("BytesSliceReply", t, func() {
		var data []byte = []byte("123.33")
		_, err := BytesSliceReply([]interface{}{data}, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestStringStringMapReply(t *testing.T) {
	convey.Convey("StringStringMapReply", t, func() {
		var key []byte = []byte("123")
		var value []byte = []byte("123.33")
		_, err := StringStringMapReply([]interface{}{key, value}, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestStringIntMapReply(t *testing.T) {
	convey.Convey("StringIntMapReply", t, func() {
		var key []byte = []byte("123")
		var value []byte = []byte("123")
		_, err := StringIntMapReply([]interface{}{key, value}, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestStringInt64MapReply(t *testing.T) {
	convey.Convey("StringInt64MapReply", t, func() {
		var key []byte = []byte("123")
		var value []byte = []byte("123")
		_, err := StringInt64MapReply([]interface{}{key, value}, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestStringBoolMapReply(t *testing.T) {
	convey.Convey("StringBoolMapReply", t, func() {
		convey.Convey("StringBoolMapReplyBranch1", func() {
			var key []byte = []byte("123")
			var value []byte = []byte("1")
			_, err := StringBoolMapReply([]interface{}{key, value}, nil)
			convey.So(err, convey.ShouldEqual, nil)
		})
		convey.Convey("StringBoolMapReplyBranch2", func() {
			var key []byte = []byte("123")
			var value []byte = []byte("1")
			_, err := StringBoolMapReply([]interface{}{key, value, 1, 1}, nil)
			convey.So(err, convey.ShouldNotEqual, nil)
		})
		convey.Convey("StringBoolMapReplyBranch3", func() {
			var key []byte = []byte("123")
			_, err := StringBoolMapReply([]interface{}{key, 1}, nil)
			convey.So(err, convey.ShouldNotEqual, nil)
		})
		convey.Convey("StringBoolMapReplyBranch4", func() {
			var key []byte = []byte("123")
			var value []byte = []byte("1")
			_, err := StringBoolMapReply([]interface{}{key, value, 1}, ErrNil)
			convey.So(err, convey.ShouldNotEqual, nil)
		})
		convey.Convey("StringBoolMapReplyBranch5", func() {
			var key []byte = []byte("123")
			_, err := StringBoolMapReply([]interface{}{key}, ErrNil)
			convey.So(err, convey.ShouldNotEqual, nil)
		})
	})
}

func TestDurationReply(t *testing.T) {
	convey.Convey("DurationReply", t, func() {
		convey.Convey("DurationReplyBranch1", func() {
			var key []byte = []byte("123")
			var value []byte = []byte("1")
			_, err := DurationReply(0, []interface{}{key, value}, nil)
			convey.So(err, convey.ShouldNotEqual, nil)
		})

		convey.Convey("DurationReplyBranch2", func() {
			_, err := DurationReply(0, int64(1), nil)
			convey.So(err, convey.ShouldEqual, nil)
		})

		convey.Convey("DurationReplyBranch3", func() {
			_, err := DurationReply(0, int64(-1), nil)
			convey.So(err, convey.ShouldEqual, nil)
		})
	})
}

func TestArgs(t *testing.T) {
	convey.Convey("args", t, func() {
		convey.Convey("DurationReplyBranch1", func() {
			s := Sort{
				By:     "1",
				Offset: 1,
				Get:    []string{"1"},
				Order:  "desc",
				Alpha:  true,
			}
			r := s.args("0")
			convey.So(len(r), convey.ShouldEqual, 10)
		})
	})
}

func TestZSliceReply(t *testing.T) {
	convey.Convey("ZSliceReply", t, func() {
		convey.Convey("ZSliceReplyBranch1", func() {
			_, err := ZSliceReply(nil, ErrNil)
			convey.So(err, convey.ShouldEqual, ErrNil)
		})

		convey.Convey("ZSliceReplyBranch2", func() {
			_, err := ZSliceReply([]interface{}{1}, nil)
			convey.So(err, convey.ShouldNotEqual, ErrNil)
		})

		convey.Convey("ZSliceReplyBranch3", func() {
			_, err := ZSliceReply([]interface{}{1, 2}, nil)
			convey.So(err, convey.ShouldNotEqual, nil)
		})

		convey.Convey("ZSliceReplyBranch4", func() {
			_, err := ZSliceReply([]interface{}{"1", 2}, nil)
			convey.So(err, convey.ShouldNotEqual, nil)
		})

		convey.Convey("ZSliceReplyBranch5", func() {
			var data []byte = []byte("1")
			_, err := ZSliceReply([]interface{}{"1", data}, nil)
			convey.So(err, convey.ShouldEqual, nil)
		})
	})
}
