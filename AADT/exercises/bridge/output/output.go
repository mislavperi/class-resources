package output

import "fmt"

type Shape interface {
	Draw() string
}

type Output struct {
	outputType string
}

func NewOutput(ot string) *Output {
	return &Output{
		outputType: ot,
	}
}

func (o *Output) DrawShape(s Shape) {
	fmt.Printf("Drawing on %s shape %s", o.outputType, s.Draw())
}
