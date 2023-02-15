package constant

type Constant struct {
	Num float32
}

func NewConstant(n float32) *Constant {
	return &Constant{
		Num: n,
	}
} 