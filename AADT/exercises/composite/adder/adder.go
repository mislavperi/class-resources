package adder

import (
	"example.com/composite/calculate"
	"example.com/composite/constant"
)

type Calc interface {
	Calcul(constants []constant.Constant) float32
}

func Add(n []constant.Constant) constant.Constant {
	c := calculate.NewCalc()
	return c.Calcul("+", n)
}
