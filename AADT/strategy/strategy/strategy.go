package strategy

import (
	"example.com/strategy/student"
)

type Algo interface {
	Sort(c []student.Student, field string)
}
