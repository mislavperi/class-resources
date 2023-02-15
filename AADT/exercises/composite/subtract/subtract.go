package subtract

import (
	"example.com/composite/calculate"
	"example.com/composite/constant"
)

func Subtract(cst []constant.Constant) constant.Constant {
	c := calculate.NewCalc()
	return c.Calcul("-", cst)
}
