package laptop

type Laptop struct {
}

func NewLaptop() *Laptop {
	return &Laptop{}
}

func (l *Laptop) GetName() string {
	return "MBP"
}
