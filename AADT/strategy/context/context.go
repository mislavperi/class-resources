package context

import (
	"example.com/strategy/strategy"
	"example.com/strategy/student"
)

type Context struct {
	Students []student.Student
	sort     strategy.Algo
}

func NewContext(s []student.Student) *Context {
	return &Context{
		Students: s,
	}
}

func (c *Context) SetSort(s strategy.Algo) {
	c.sort = s
}

func (c *Context) ApplySort(s []student.Student, field string) {
	c.sort.Sort(c.Students, field)
}
