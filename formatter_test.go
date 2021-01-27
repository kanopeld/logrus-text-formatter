package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTextFormatter_Format(t *testing.T) {
	convey.Convey("Testing time formatting", t, func() {
		f := &TextFormatter{
			TimeFormat: time.RFC3339,
		}

		t := time.Now()
		logEntity := &logrus.Entry{
			Time: t,
		}

		resB, err := f.Format(logEntity)
		convey.So(err, convey.ShouldBeNil)
		resS := string(resB)
		convey.So(resS, convey.ShouldEqual, fmt.Sprintf("[panic]: %s\n", t.Format(time.RFC3339)))
	})
}
