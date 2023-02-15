package insert

import (
	"reflect"

	"example.com/strategy/student"
)

type Insert struct {
}

func NewInsert() *Insert {
	return &Insert{}
}

func (i *Insert) Sort(c []student.Student, field string) {
	for g := 1; g < len(c); g++ {
		j := g
		for j >= 1 && getField(c[j], field) < getField(c[j-1], field) {
			c[j], c[j-1] = c[j-1], c[j]
			j--
		}
	}
}

func getField(v student.Student, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}
