package redisx

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"

	"git.code.oa.com/NGTest/gomonkey"
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"
)

func TestZadd(t *testing.T) {
	convey.Convey("Zadd", t, func() {
		convey.Convey("Zadd", func() {
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
			members := &Z{
				Score:  2.2,
				Member: "2.2",
			}
			rsp, err := client.ZAdd(ctx, "1", members)
			convey.So(rsp, convey.ShouldEqual, 12)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZAddNX(t *testing.T) {
	convey.Convey("ZAddNX", t, func() {
		convey.Convey("ZAddNX", func() {
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
			members := &Z{
				Score:  2.2,
				Member: "2.2",
			}
			rsp, err := client.ZAddNX(ctx, "1", members)
			convey.So(rsp, convey.ShouldEqual, 12)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZAddXX(t *testing.T) {
	convey.Convey("ZAddXX", t, func() {
		convey.Convey("ZAddXX", func() {
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
			members := &Z{
				Score:  2.2,
				Member: "2.2",
			}
			rsp, err := client.ZAddXX(ctx, "1", members)
			convey.So(rsp, convey.ShouldEqual, 12)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZAddCh(t *testing.T) {
	convey.Convey("ZAddCh", t, func() {
		convey.Convey("ZAddCh", func() {
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
			members := &Z{
				Score:  2.2,
				Member: "2.2",
			}
			rsp, err := client.ZAddCh(ctx, "1", members)
			convey.So(rsp, convey.ShouldEqual, 12)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZAddNXCh(t *testing.T) {
	convey.Convey("ZAddNXCh", t, func() {
		convey.Convey("ZAddNXCh", func() {
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
			members := &Z{
				Score:  2.2,
				Member: "2.2",
			}
			rsp, err := client.ZAddNXCh(ctx, "1", members)
			convey.So(rsp, convey.ShouldEqual, 12)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZAddXXCh(t *testing.T) {
	convey.Convey("ZAddXXCh", t, func() {
		convey.Convey("ZAddXXCh", func() {
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
			members := &Z{
				Score:  2.2,
				Member: "2.2",
			}
			rsp, err := client.ZAddXXCh(ctx, "1", members)
			convey.So(rsp, convey.ShouldEqual, 12)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZIncr(t *testing.T) {
	convey.Convey("ZIncr", t, func() {
		convey.Convey("ZIncr", func() {
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
			members := Z{
				Score:  2.2,
				Member: "2.2",
			}
			rsp, err := client.ZIncr(ctx, "1", members)
			convey.So(rsp, convey.ShouldEqual, 2.2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZIncrNX(t *testing.T) {
	convey.Convey("ZIncrNX", t, func() {
		convey.Convey("ZIncrNX", func() {
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
			members := Z{
				Score:  2.2,
				Member: "2.2",
			}
			rsp, err := client.ZIncrNX(ctx, "1", members)
			convey.So(rsp, convey.ShouldEqual, 2.2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZIncrXX(t *testing.T) {
	convey.Convey("ZIncrXX", t, func() {
		convey.Convey("ZIncrXX", func() {
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
			members := Z{
				Score:  2.2,
				Member: "2.2",
			}
			rsp, err := client.ZIncrXX(ctx, "1", members)
			convey.So(rsp, convey.ShouldEqual, 2.2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZCard(t *testing.T) {
	convey.Convey("ZCard", t, func() {
		convey.Convey("ZCard", func() {
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
			rsp, err := client.ZCard(ctx, "1")
			convey.So(rsp, convey.ShouldEqual, 15)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZCount(t *testing.T) {
	convey.Convey("ZCount", t, func() {
		convey.Convey("ZCount", func() {
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
			rsp, err := client.ZCount(ctx, "1", "0", "0")
			convey.So(rsp, convey.ShouldEqual, 15)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZLexCount(t *testing.T) {
	convey.Convey("ZLexCount", t, func() {
		convey.Convey("ZLexCount", func() {
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
			rsp, err := client.ZLexCount(ctx, "1", "0", "0")
			convey.So(rsp, convey.ShouldEqual, 15)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZIncrBy(t *testing.T) {
	convey.Convey("ZIncrBy", t, func() {
		convey.Convey("ZIncrBy", func() {
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
			rsp, err := client.ZIncrBy(ctx, "1", 0, "0")
			convey.So(rsp, convey.ShouldEqual, 2.2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZMScore(t *testing.T) {
	convey.Convey("ZMScore", t, func() {
		convey.Convey("ZMScore", func() {
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
			rsp, err := client.ZMScore(ctx, "1", "0")
			convey.So(len(rsp), convey.ShouldEqual, 4)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZPopMax(t *testing.T) {
	convey.Convey("ZPopMax", t, func() {
		convey.Convey("ZPopMax", func() {
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
			rsp, err := client.ZPopMax(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZPopMin(t *testing.T) {
	convey.Convey("ZPopMin", t, func() {
		convey.Convey("ZPopMin", func() {
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
			rsp, err := client.ZPopMin(ctx, "1")
			convey.So(len(rsp), convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRange(t *testing.T) {
	convey.Convey("ZRange", t, func() {
		convey.Convey("ZRange", func() {
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
			rsp, err := client.ZRange(ctx, "1", 0, 0)
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRangeWithScores(t *testing.T) {
	convey.Convey("ZRangeWithScores", t, func() {
		convey.Convey("ZRangeWithScores", func() {
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
			rsp, err := client.ZRangeWithScores(ctx, "1", 0, 0)
			convey.So(len(rsp), convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRangeByScore(t *testing.T) {
	convey.Convey("ZRangeByScore", t, func() {
		convey.Convey("ZRangeByScore", func() {
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
			var opt ZRangeBy
			rsp, err := client.ZRangeByScore(ctx, "1", opt)
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRangeByScoreWithScores(t *testing.T) {
	convey.Convey("ZRangeByScoreWithScores", t, func() {
		convey.Convey("ZRangeByScoreWithScores", func() {
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
			var opt ZRangeBy
			rsp, err := client.ZRangeByScoreWithScores(ctx, "1", opt)
			convey.So(len(rsp), convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRank(t *testing.T) {
	convey.Convey("ZRank", t, func() {
		convey.Convey("ZRank", func() {
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
			rsp, err := client.ZRank(ctx, "1", "2")
			convey.So(rsp, convey.ShouldEqual, 16)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRem(t *testing.T) {
	convey.Convey("ZRem", t, func() {
		convey.Convey("ZRem", func() {
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
			rsp, err := client.ZRem(ctx, "1", "2")
			convey.So(rsp, convey.ShouldEqual, 16)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRemRangeByRank(t *testing.T) {
	convey.Convey("ZRemRangeByRank", t, func() {
		convey.Convey("ZRemRangeByRank", func() {
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
			rsp, err := client.ZRemRangeByRank(ctx, "1", 0, 0)
			convey.So(rsp, convey.ShouldEqual, 18)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRemRangeByScore(t *testing.T) {
	convey.Convey("ZRemRangeByScore", t, func() {
		convey.Convey("ZRemRangeByScore", func() {
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
			rsp, err := client.ZRemRangeByScore(ctx, "1", "0", "0")
			convey.So(rsp, convey.ShouldEqual, 19)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRevRange(t *testing.T) {
	convey.Convey("ZRevRange", t, func() {
		convey.Convey("ZRevRange", func() {
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
			rsp, err := client.ZRevRange(ctx, "1", 0, 0)
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRevRangeWithScores(t *testing.T) {
	convey.Convey("ZRevRangeWithScores", t, func() {
		convey.Convey("ZRevRangeWithScores", func() {
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
			rsp, err := client.ZRevRangeWithScores(ctx, "1", 0, 0)
			convey.So(len(rsp), convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRevRangeByScore(t *testing.T) {
	convey.Convey("ZRevRangeByScore", t, func() {
		convey.Convey("ZRevRangeByScore", func() {
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
			var opt ZRangeBy
			rsp, err := client.ZRevRangeByScore(ctx, "1", opt)
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRevRangeByScoreWithScores(t *testing.T) {
	convey.Convey("ZRevRangeByScoreWithScores", t, func() {
		convey.Convey("ZRevRangeByScoreWithScores", func() {
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
			var opt ZRangeBy
			rsp, err := client.ZRevRangeByScoreWithScores(ctx, "1", opt)
			convey.So(len(rsp), convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRevRank(t *testing.T) {
	convey.Convey("ZRevRank", t, func() {
		convey.Convey("ZRevRank", func() {
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
			rsp, err := client.ZRevRank(ctx, "1", "2")
			convey.So(rsp, convey.ShouldEqual, 16)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZScore(t *testing.T) {
	convey.Convey("ZScore", t, func() {
		convey.Convey("ZScore", func() {
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
			rsp, err := client.ZScore(ctx, "1", "2")
			convey.So(rsp, convey.ShouldEqual, 2.2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}

func TestZRandMember(t *testing.T) {
	convey.Convey("ZRandMember", t, func() {
		convey.Convey("ZRandMember", func() {
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
			rsp, err := client.ZRandMember(ctx, "1", 0, false)
			convey.So(len(rsp), convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldBeNil)

		})
	})
}
