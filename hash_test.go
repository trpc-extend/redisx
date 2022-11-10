package redisx

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"

	"git.code.oa.com/NGTest/gomonkey"
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"
)

func TestHSet(t *testing.T) {
	convey.Convey("HSet", t, func() {
		convey.Convey("HSetDefault", func() {
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
			rsp, err := client.HSet(ctx, "1", 2)
			convey.So(rsp, convey.ShouldEqual, 7)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHSetNX(t *testing.T) {
	convey.Convey("HSetNX", t, func() {
		convey.Convey("HSetNX", func() {
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
			rsp, err := client.HSetNX(ctx, "1", "act", 2)
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHGet(t *testing.T) {
	convey.Convey("HGet", t, func() {
		convey.Convey("HGet", func() {
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
			rsp, err := client.HGet(ctx, "1", "act")
			convey.So(rsp, convey.ShouldEqual, "111")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHExists(t *testing.T) {
	convey.Convey("HExists", t, func() {
		convey.Convey("HExists", func() {
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
			rsp, err := client.HExists(ctx, "1", "act")
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHDel(t *testing.T) {
	convey.Convey("HDel", t, func() {
		convey.Convey("HDel", func() {
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
			rsp, err := client.HDel(ctx, "1", "act")
			convey.So(rsp, convey.ShouldEqual, 8)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHLen(t *testing.T) {
	convey.Convey("HLen", t, func() {
		convey.Convey("HLen", func() {
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
			rsp, err := client.HLen(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 9)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHStrLen(t *testing.T) {
	convey.Convey("HStrLen", t, func() {
		convey.Convey("HStrLen", func() {
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
			rsp, err := client.HStrLen(ctx, "1", "zz")
			convey.So(rsp, convey.ShouldEqual, 10)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHIncrBy(t *testing.T) {
	convey.Convey("HIncrBy", t, func() {
		convey.Convey("HIncrBy", func() {
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
			rsp, err := client.HIncrBy(ctx, "1", "zz", 66)
			convey.So(rsp, convey.ShouldEqual, 11)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHIncrByFloat(t *testing.T) {
	convey.Convey("HIncrByFloat", t, func() {
		convey.Convey("HIncrByFloat", func() {
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
			rsp, err := client.HIncrByFloat(ctx, "1", "zz", 66)
			convey.So(rsp, convey.ShouldEqual, 2.2)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHMSet(t *testing.T) {
	convey.Convey("HMSet", t, func() {
		convey.Convey("HMSet", func() {
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
			rsp, err := client.HMSet(ctx, "1", "zz", 66)
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHMGet(t *testing.T) {
	convey.Convey("HMGet", t, func() {
		convey.Convey("HMGet", func() {
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
			rsp, err := client.HMGet(ctx, "1", "zz")
			convey.So(len(rsp), convey.ShouldEqual, 4)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHKeys(t *testing.T) {
	convey.Convey("HKeys", t, func() {
		convey.Convey("HKeys", func() {
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
			rsp, err := client.HKeys(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 5)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHVals(t *testing.T) {
	convey.Convey("HVals", t, func() {
		convey.Convey("HVals", func() {
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
			rsp, err := client.HVals(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 5)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHGetAll(t *testing.T) {
	convey.Convey("HGetAll", t, func() {
		convey.Convey("HGetAll", func() {
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
			rsp, err := client.HGetAll(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
