package divider

import (
	"example.com/composite/calculate"
	"example.com/composite/constant"
)

type Calc interface {
	Calcul(op string, constants []constant.Constant) float32
}

func Divide(consts []constant.Constant) constant.Constant {
	c := calculate.NewCalc()
	return c.Calcul("/", consts)
	
}
