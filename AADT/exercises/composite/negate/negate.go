package negate

import "example.com/composite/constant"


func NegateN(n constant.Constant) constant.Constant {
	n.Num = -n.Num
	return n
}