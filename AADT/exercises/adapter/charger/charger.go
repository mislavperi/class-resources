package charger

import "fmt"

type Laptop interface {
	GetName() string
}

type EuSocket interface {
	ReturnSocket() string
}

type Charger struct {
}

func NewCharger() *Charger {
	return &Charger{}
}

func (c *Charger) Charge(laptop Laptop, euSocket EuSocket) {
	laptopName := laptop.GetName()
	socketName := euSocket.ReturnSocket()
	res := fmt.Sprintf("Charging %s with %s", laptopName, socketName)
	fmt.Println(res)
}
