package calculate

import (
	"example.com/composite/constant"
)

type Calc struct {
}

func NewCalc() *Calc {
	return &Calc{}
}

func (c *Calc) Calcul(op string, constants []constant.Constant) constant.Constant {
	switch op {
	case "+":
		return *constant.NewConstant(constants[0].Num + constants[1].Num)
	case "*":
		return *constant.NewConstant(constants[0].Num * constants[1].Num)
	case "/":
		return *constant.NewConstant(constants[0].Num / constants[1].Num)
	case "-":
		return *constant.NewConstant(constants[0].Num - constants[1].Num)
	default:
		return *constant.NewConstant(0)
	}

}
