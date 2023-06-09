package testbdd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("when add the two numbers", func() {
			c := a + b

			Convey("then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
