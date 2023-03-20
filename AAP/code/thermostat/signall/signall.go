package signall

import "fmt"

type SingalL struct {
}

func NewSignalL() *SingalL {
	return &SingalL{}
}

func (s *SingalL) TurnOn() {
	fmt.Println("turned on")

}
