package fibonacci_test

import (
	"testing"

	"example.com/unittests/fibonacci"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFibonacci(t *testing.T) {
	Convey("Given the Fibonacci sequance generation function", t, func() {
		Convey("For number 5", func() {
			result := fibonacci.Generate(5)
			Convey("Result should be 5", func() {
				So(result, ShouldEqual, 5)
			})
		})
		Convey("For number 11", func() {
			result := fibonacci.Generate(11)
			Convey("Result should be 89", func() {
				So(result, ShouldEqual, 89)
			})
		})
		Convey("For number 2", func() {
			result := fibonacci.Generate(2)
			Convey("Result should be 1", func() {
				So(result, ShouldEqual, 1)
			})
		})
		Convey("For number 18", func() {
			result := fibonacci.Generate(18)
			Convey("Result should be 5", func() {
				So(result, ShouldEqual, 2584)
			})
		})
	})
}
