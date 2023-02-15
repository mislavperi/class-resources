package shape

type Shape struct {
	shapeType string
}

func NewShape(sType string) *Shape {
	return &Shape{
		shapeType: sType,
	}
}

func (s *Shape) Draw() string {
	return s.shapeType
}
