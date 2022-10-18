package adapters

import "fmt"

type GbAdapter struct {
	gbSocket GbSocket
	euSocket EuSocket
}

type EuSocket interface {

	ReturnSocket() string
}

type GbSocket interface {
	GiveGbCurrent() string
}

func NewGbAdapter(gbSocket GbSocket, euSocket EuSocket) *GbAdapter {
	return &GbAdapter{
		gbSocket: gbSocket,
		euSocket: euSocket,
	}
}

func (e *GbAdapter) AdaptToEuPower() EuSocket {
	fmt.Println("adapting from British to EU power")
	return e.euSocket
}
