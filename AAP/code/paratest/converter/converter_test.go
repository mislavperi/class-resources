package converter_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/paratest/converter"
	. "github.com/smartystreets/goconvey/convey"
)

func parametrize[V any, T any](fn T, allValues [][]V) {
	v := reflect.ValueOf(fn)
	for _, a := range allValues {
		vargs := make([]reflect.Value, len(a))

		for i, b := range a {
			vargs[i] = reflect.ValueOf(b)
		}
		v.Call(vargs)
	}
}

func TestConverter(t *testing.T) {
	Convey("Given the two converters", t, func() {
		c := converter.ConverterC{}
		f := converter.ConverterF{}
		Convey("Given the Celsius converter", func() {
			testArgs := [][]any{
				{t, 1, -17}, {t, 3, -16},
			}
			test := func(t *testing.T, input int, expected int) {
				assert.Equal(t, c.Convert(input), expected)
			}
			parametrize(test, testArgs)
		})
		Convey("Given the Fahrenheit converter", func() {
			testArgs := [][]any{
				{t, 1, 33}, {t, 3, 37},
			}
			test := func(t *testing.T, input int, expected int) {
				assert.Equal(t, f.Convert(input), expected)
			}
			parametrize(test, testArgs)
		})
	})
}
