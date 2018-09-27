package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStrFirstToUpper(t *testing.T) {

	Convey("首字母大写", t, func() {
		So(strFirstToUpper("hello"), ShouldEqual, "Hello")
	})
}