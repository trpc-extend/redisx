package redisx

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestGetInstance(t *testing.T) {
	convey.Convey("getInstance", t, func() {
		convey.Convey("getInstance", func() {
			getInstance()
		})
	})
}

func TestCommReport(t *testing.T) {
	convey.Convey("commReport", t, func() {
		convey.Convey("commReport", func() {
			commReport(nil, nil)
		})
	})
}

func TestCommReportWithSuffix(t *testing.T) {
	convey.Convey("commReportWithSuffix", t, func() {
		convey.Convey("commReportWithSuffix", func() {
			commReportWithSuffix("diy", nil, nil)
		})
	})
}

func TestCmdReport(t *testing.T) {
	convey.Convey("cmdReport", t, func() {
		convey.Convey("cmdReport", func() {
			var t time.Time
			cmdReport("diy", "service", "get", t, nil)
		})
	})
}
