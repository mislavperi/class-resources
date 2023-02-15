package checkup

import "fmt"

type Checkup struct {
}

func NewCheckUp() *Checkup {
	return &Checkup{}
}

func (c *Checkup) DoCheckup() {
	fmt.Println("Doing basic checkup")
}
