package main

import (
	"fmt"

	"example.com/composite/adder"
	"example.com/composite/constant"
	"example.com/composite/divider"
	"example.com/composite/multiply"
	"example.com/composite/negate"
	"example.com/composite/subtract"
)

func main() {
	fmt.Println(
		divider.Divide(
			[]constant.Constant{
				multiply.Multi(
					[]constant.Constant{
						negate.NegateN(
							*constant.NewConstant(5),
						),
						divider.Divide(
							[]constant.Constant{
								*constant.NewConstant(9),
								*constant.NewConstant(6),
							},
						),
					},
				),
				adder.Add(
					[]constant.Constant{
						*constant.NewConstant(7),
						subtract.Subtract(
							[]constant.Constant{
								*constant.NewConstant(2),
								*constant.NewConstant(1.5),
							},
						),
					},
				),
			},
		),
	)
}
