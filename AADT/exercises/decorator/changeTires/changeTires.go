package changetires

import "fmt"

type Checkup interface {
	DoCheckup()
}

type ChangeTires struct {
	Checkup Checkup
}

func NewChangeTire(c Checkup) *ChangeTires {
	return &ChangeTires{
		Checkup: c,
	}
}

func (c *ChangeTires) ChangeTires() {
	fmt.Println("Changing tires")
	c.Checkup.DoCheckup()
}
