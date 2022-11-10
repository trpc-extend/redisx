package redisx

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"

	"git.code.oa.com/NGTest/gomonkey"
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"
)

func TestSAdd(t *testing.T) {
	convey.Convey("SAdd", t, func() {
		convey.Convey("SAdd", func() {
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
			rsp, err := client.SAdd(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 27)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSCard(t *testing.T) {
	convey.Convey("SCard", t, func() {
		convey.Convey("SCard", func() {
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
			rsp, err := client.SCard(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 28)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSDiff(t *testing.T) {
	convey.Convey("SDiff", t, func() {
		convey.Convey("SDiff", func() {
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
			rsp, err := client.SDiff(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 3)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSDiffStore(t *testing.T) {
	convey.Convey("SDiffStore", t, func() {
		convey.Convey("SDiffStore", func() {
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
			rsp, err := client.SDiffStore(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 29)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSInter(t *testing.T) {
	convey.Convey("SInter", t, func() {
		convey.Convey("SInter", func() {
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
			rsp, err := client.SInter(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 3)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSInterStore(t *testing.T) {
	convey.Convey("SInterStore", t, func() {
		convey.Convey("SInterStore", func() {
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
			rsp, err := client.SInterStore(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 30)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSIsMember(t *testing.T) {
	convey.Convey("SIsMember", t, func() {
		convey.Convey("SIsMember", func() {
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
			rsp, err := client.SIsMember(ctx, "1", "zz")
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSMembers(t *testing.T) {
	convey.Convey("SMembers", t, func() {
		convey.Convey("SMembers", func() {
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
			rsp, err := client.SMembers(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 3)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSMove(t *testing.T) {
	convey.Convey("SMove", t, func() {
		convey.Convey("SMove", func() {
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
			rsp, err := client.SMove(ctx, "1", "2", 3)
			convey.So(rsp, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSPop(t *testing.T) {
	convey.Convey("SPop", t, func() {
		convey.Convey("SPop", func() {
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
			rsp, err := client.SPop(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, "1")
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSPopN(t *testing.T) {
	convey.Convey("SPopN", t, func() {
		convey.Convey("SPopN", func() {
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
			rsp, err := client.SPopN(ctx, "1", 1)
			convey.So(len(rsp), convey.ShouldEqual, 0)
			convey.So(err, convey.ShouldNotBeNil)

		})
	})
}

func TestSRandMember(t *testing.T) {
	convey.Convey("SRandMember", t, func() {
		convey.Convey("SRandMember", func() {
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
			rsp, err := client.SRandMember(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSRandMemberN(t *testing.T) {
	convey.Convey("SRandMemberN", t, func() {
		convey.Convey("SRandMemberN", func() {
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
			rsp, err := client.SRandMemberN(ctx, "1", 1)
			convey.So(len(rsp), convey.ShouldEqual, 0)
			convey.So(err, convey.ShouldNotBeNil)

		})
	})
}

func TestSRem(t *testing.T) {
	convey.Convey("SRem", t, func() {
		convey.Convey("SRem", func() {
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
			rsp, err := client.SRem(ctx, "1", 1)
			convey.So(rsp, convey.ShouldEqual, 31)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSUnion(t *testing.T) {
	convey.Convey("SUnion", t, func() {
		convey.Convey("SUnion", func() {
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
			rsp, err := client.SUnion(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 3)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestSUnionStore(t *testing.T) {
	convey.Convey("SUnionStore", t, func() {
		convey.Convey("SUnionStore", func() {
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
			rsp, err := client.SUnionStore(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 31)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}
