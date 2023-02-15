package main

import (
	"fmt"

	"example.com/strategy/bubble"
	"example.com/strategy/context"
	"example.com/strategy/student"
)

type Student struct {
	Grade int
	Name  string
	LName string
}

func main() {
	st := []student.Student{
		{
			Grade: 3,
			Name:  "z",
			LName: "PErić",
		},
		{
			Grade: 5,
			Name:  "c",
			LName: "asdasdas",
		},
		{
			Grade: 1,
			Name:  "g",
			LName: "qwrqrq",
		},
		{
			Grade: 2,
			Name:  "l",
			LName: "qwrqwrq",
		},
	}
	ctx := context.NewContext(st)
	bubble := bubble.NewBubble()
	ctx.SetSort(bubble)
	ctx.ApplySort(ctx.Students, "Grade")
	fmt.Println(ctx)
}
