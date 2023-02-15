package carwash

import "fmt"

type Checkup interface {
	DoCheckup()
}

type CarWash struct {
	Checkup Checkup
}

func NewCarWash(c Checkup) *CarWash {
	return &CarWash{
		Checkup: c,
	}
}

func (c *CarWash) WashCar() {
	fmt.Println("Wash car")
	c.Checkup.DoCheckup()
}
