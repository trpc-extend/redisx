package redisx

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"

	"git.code.oa.com/NGTest/gomonkey"
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"
)

func TestLPush(t *testing.T) {
	convey.Convey("LPush", t, func() {
		convey.Convey("LPush", func() {
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
			rsp, err := client.LPush(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 20)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLPushX(t *testing.T) {
	convey.Convey("LPushX", t, func() {
		convey.Convey("LPushX", func() {
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
			rsp, err := client.LPushX(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 21)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLRange(t *testing.T) {
	convey.Convey("LRange", t, func() {
		convey.Convey("LRange", func() {
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
			rsp, err := client.LRange(ctx, "1", 0, 0)
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLRem(t *testing.T) {
	convey.Convey("LRem", t, func() {
		convey.Convey("LRem", func() {
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
			rsp, err := client.LRem(ctx, "1", 0, 0)
			convey.So(rsp, convey.ShouldEqual, 22)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLSet(t *testing.T) {
	convey.Convey("LSet", t, func() {
		convey.Convey("LSet", func() {
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
			err := client.LSet(ctx, "1", 0, 0)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLTrim(t *testing.T) {
	convey.Convey("LTrim", t, func() {
		convey.Convey("LTrim", func() {
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
			err := client.LTrim(ctx, "1", 0, 0)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestRPop(t *testing.T) {
	convey.Convey("RPop", t, func() {
		convey.Convey("RPop", func() {
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
			rsp, err := client.RPop(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, "1")
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestRPopLPush(t *testing.T) {
	convey.Convey("RPopLPush", t, func() {
		convey.Convey("RPopLPush", func() {
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
			rsp, err := client.RPopLPush(ctx, "1", "2")
			convey.So(rsp, convey.ShouldEqual, "2")
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestRPush(t *testing.T) {
	convey.Convey("RPush", t, func() {
		convey.Convey("RPush", func() {
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
			rsp, err := client.RPush(ctx, "1", "2")
			convey.So(rsp, convey.ShouldEqual, 23)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestRPushX(t *testing.T) {
	convey.Convey("RPushX", t, func() {
		convey.Convey("RPushX", func() {
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
			rsp, err := client.RPushX(ctx, "1", "2")
			convey.So(rsp, convey.ShouldEqual, 24)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestBLPop(t *testing.T) {
	convey.Convey("BLPop", t, func() {
		convey.Convey("BLPop", func() {
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
			var tt time.Duration
			rsp, err := client.BLPop(ctx, tt, "2")
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestBRPop(t *testing.T) {
	convey.Convey("BRPop", t, func() {
		convey.Convey("BRPop", func() {
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
			var tt time.Duration
			rsp, err := client.BRPop(ctx, tt, "2")
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestBRPopLPush(t *testing.T) {
	convey.Convey("BRPopLPush", t, func() {
		convey.Convey("BRPopLPush", func() {
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
			var tt time.Duration
			rsp, err := client.BRPopLPush(ctx, "source", "dst", tt)
			convey.So(rsp, convey.ShouldEqual, "1")
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLIndex(t *testing.T) {
	convey.Convey("LIndex", t, func() {
		convey.Convey("LIndex", func() {
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
			rsp, err := client.LIndex(ctx, "key", 0)
			convey.So(rsp, convey.ShouldEqual, "1")
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLInsert(t *testing.T) {
	convey.Convey("LInsert", t, func() {
		convey.Convey("LInsert", func() {
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
			rsp, err := client.LInsert(ctx, "key", "op", 1, 2)
			convey.So(rsp, convey.ShouldEqual, 25)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLInsertBefore(t *testing.T) {
	convey.Convey("LInsertBefore", t, func() {
		convey.Convey("LInsertBefore", func() {
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
			rsp, err := client.LInsertBefore(ctx, "key", 1, 2)
			convey.So(rsp, convey.ShouldEqual, 25)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLInsertAfter(t *testing.T) {
	convey.Convey("LInsertAfter", t, func() {
		convey.Convey("LInsertAfter", func() {
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
			rsp, err := client.LInsertAfter(ctx, "key", 1, 2)
			convey.So(rsp, convey.ShouldEqual, 25)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLLen(t *testing.T) {
	convey.Convey("LLen", t, func() {
		convey.Convey("LLen", func() {
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
			rsp, err := client.LLen(ctx, "key")
			convey.So(rsp, convey.ShouldEqual, 26)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestLPop(t *testing.T) {
	convey.Convey("LPop", t, func() {
		convey.Convey("LPop", func() {
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
			rsp, err := client.LPop(ctx, "key")
			convey.So(rsp, convey.ShouldEqual, "3")
			convey.So(err, convey.ShouldBeNil)

		})
	})
}
