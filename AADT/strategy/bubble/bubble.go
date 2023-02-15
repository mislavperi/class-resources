package bubble

import (
	"reflect"

	"example.com/strategy/student"
)

type Bubble struct {
}

func NewBubble() *Bubble {
	return &Bubble{}
}

func (b *Bubble) Sort(c []student.Student, field string) {
	for i := 0; i < len(c)-1; i++ {
		for j := 0; j < len(c)-i-1; j++ {
			if getField(c[j], field) > getField(c[j+1], field) {
				c[j], c[j+1] = c[j+1], c[j]
			}
		}
	}
}

func getField(v student.Student, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}
