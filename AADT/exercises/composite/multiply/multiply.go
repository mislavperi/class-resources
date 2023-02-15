package multiply

import (
	"example.com/composite/calculate"
	"example.com/composite/constant"
)

func Multi(constat []constant.Constant) constant.Constant {
	c := calculate.NewCalc()
	return c.Calcul("*", constat)
}
