package redisx

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"

	"git.code.oa.com/NGTest/gomonkey"
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"
)

func TestSet(t *testing.T) {
	convey.Convey("Set", t, func() {
		convey.Convey("SetDefault", func() {
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
			err := client.Set(ctx, "1", 2, 0)
			convey.So(err, convey.ShouldBeNil)

			err = client.Set(ctx, "1", 2, -1)
			convey.So(err, convey.ShouldBeNil)

			err = client.Set(ctx, "1", 2, -2)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestSetNX(t *testing.T) {
	convey.Convey("SetNX", t, func() {
		convey.Convey("SetNX", func() {
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
			_, err := client.SetNX(ctx, "1", 2, 0)
			convey.So(err, convey.ShouldBeNil)

			_, err = client.SetNX(ctx, "1", 2, -1)
			convey.So(err, convey.ShouldBeNil)

			_, err = client.SetNX(ctx, "1", 2, -2)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestSetEX(t *testing.T) {
	convey.Convey("SetEX", t, func() {
		convey.Convey("SetEX", func() {
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
			_, err := client.SetNX(ctx, "1", 2, 0)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestPSetEX(t *testing.T) {
	convey.Convey("PSetEX", t, func() {
		convey.Convey("PSetEX", func() {
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
			err := client.PSetEX(ctx, "1", 2, 0)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestSetXX(t *testing.T) {
	convey.Convey("SetXX", t, func() {
		convey.Convey("SetXX", func() {
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
			_, err := client.SetXX(ctx, "1", 2, 0)
			convey.So(err, convey.ShouldBeNil)

			_, err = client.SetXX(ctx, "1", 2, -1)
			convey.So(err, convey.ShouldBeNil)

			_, err = client.SetXX(ctx, "1", 2, -2)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestSetRange(t *testing.T) {
	convey.Convey("SetRange", t, func() {
		convey.Convey("SetRange", func() {
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
			_, err := client.SetRange(ctx, "1", 2, "0")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestGet(t *testing.T) {
	convey.Convey("Get", t, func() {
		convey.Convey("Get", func() {
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
			rsp, err := client.Get(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, "test")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestGetRange(t *testing.T) {
	convey.Convey("GetRange", t, func() {
		convey.Convey("GetRange", func() {
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
			rsp, err := client.GetRange(ctx, "1", 0, 2)
			convey.So(rsp, convey.ShouldEqual, "test")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestGetSet(t *testing.T) {
	convey.Convey("GetSet", t, func() {
		convey.Convey("GetSet", func() {
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
			rsp, err := client.GetSet(ctx, "1", 0)
			convey.So(rsp, convey.ShouldEqual, "test")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestIncr(t *testing.T) {
	convey.Convey("Incr", t, func() {
		convey.Convey("Incr", func() {
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
			rsp, err := client.Incr(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestIncrBy(t *testing.T) {
	convey.Convey("IncrBy", t, func() {
		convey.Convey("IncrBy", func() {
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
			rsp, err := client.IncrBy(ctx, "1", 2)
			convey.So(rsp, convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestIncrByFloat(t *testing.T) {
	convey.Convey("IncrByFloat", t, func() {
		convey.Convey("IncrByFloat", func() {
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
			rsp, err := client.IncrByFloat(ctx, "1", 2.2)
			convey.So(rsp, convey.ShouldEqual, 2.2)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestDecr(t *testing.T) {
	convey.Convey("Decr", t, func() {
		convey.Convey("Decr", func() {
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
			rsp, err := client.Decr(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 3)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestDecrBy(t *testing.T) {
	convey.Convey("DecrBy", t, func() {
		convey.Convey("DecrBy", func() {
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
			rsp, err := client.DecrBy(ctx, "1", -1)
			convey.So(rsp, convey.ShouldEqual, 4)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestStrLen(t *testing.T) {
	convey.Convey("StrLen", t, func() {
		convey.Convey("StrLen", func() {
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
			rsp, err := client.StrLen(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 5)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestAppend(t *testing.T) {
	convey.Convey("Append", t, func() {
		convey.Convey("Append", func() {
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
			rsp, err := client.Append(ctx, "1", "2")
			convey.So(rsp, convey.ShouldEqual, 6)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestMGet(t *testing.T) {
	convey.Convey("MGet", t, func() {
		convey.Convey("MGet", func() {
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
			rsp, err := client.MGet(ctx, "1", "2")
			convey.So(len(rsp), convey.ShouldEqual, 3)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestMSet(t *testing.T) {
	convey.Convey("MSet", t, func() {
		convey.Convey("MSet", func() {
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
			err := client.MSet(ctx, "1", "2")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestMSetNX(t *testing.T) {
	convey.Convey("MSetNX", t, func() {
		convey.Convey("MSetNX", func() {
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
			_, err := client.MSetNX(ctx, "1", "2")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
