package redisx

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"

	"git.code.oa.com/NGTest/gomonkey"
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"
)

func TestDel(t *testing.T) {
	convey.Convey("Del", t, func() {
		convey.Convey("Del", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.Del(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 32)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestExists(t *testing.T) {
	convey.Convey("Exists", t, func() {
		convey.Convey("Exists", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.Exists(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestExpire(t *testing.T) {
	convey.Convey("Expire", t, func() {
		convey.Convey("Expire", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.Expire(ctx, "1", 0)
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestExpireAt(t *testing.T) {
	convey.Convey("ExpireAt", t, func() {
		convey.Convey("ExpireAt", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			var ti time.Time
			rsp, err := client.ExpireAt(ctx, "1", ti)
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestKeys(t *testing.T) {
	convey.Convey("Keys", t, func() {
		convey.Convey("Keys", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.Keys(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestPersist(t *testing.T) {
	convey.Convey("Persist", t, func() {
		convey.Convey("Persist", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.Persist(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestPExpire(t *testing.T) {
	convey.Convey("PExpire", t, func() {
		convey.Convey("PExpire", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.PExpire(ctx, "1", 0)
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestPExpireAt(t *testing.T) {
	convey.Convey("PExpireAt", t, func() {
		convey.Convey("PExpireAt", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			var ti time.Time
			rsp, err := client.PExpireAt(ctx, "1", ti)
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestPTTL(t *testing.T) {
	convey.Convey("PTTL", t, func() {
		convey.Convey("PTTL", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.PTTL(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 0)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestRandomKey(t *testing.T) {
	convey.Convey("RandomKey", t, func() {
		convey.Convey("RandomKey", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.RandomKey(ctx)
			convey.So(rsp, convey.ShouldEqual, "1111")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestTouch(t *testing.T) {
	convey.Convey("Touch", t, func() {
		convey.Convey("Touch", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.Touch(ctx, "11")
			convey.So(rsp, convey.ShouldEqual, 33)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestTTL(t *testing.T) {
	convey.Convey("TTL", t, func() {
		convey.Convey("TTL", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.TTL(ctx, "11")
			convey.So(rsp, convey.ShouldEqual, 0)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestType(t *testing.T) {
	convey.Convey("Type", t, func() {
		convey.Convey("Type", func() {
			var mp MyProxy
			var p1 = gomonkey.ApplyFunc(NewClient, func(name string, opts ...client.Option) Client {
				imp := &impDao{
					Client:          mp,
					MaxRetries:      0,
					MinRetryBackoff: 8 * time.Millisecond,
					MaxRetryBackoff: 16 * time.Millisecond,
				}
				return imp
			})
			defer p1.Reset()

			ctx := trpc.BackgroundContext()
			client := NewClient("test")
			rsp, err := client.Type(ctx, "11")
			convey.So(rsp, convey.ShouldEqual, "set")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
