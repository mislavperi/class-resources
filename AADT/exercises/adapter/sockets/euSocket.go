package sockets

type EuSocket struct {

}

func NewEuSocket() *EuSocket {
	return &EuSocket{

	}
}

func (e *EuSocket) ReturnSocket() string {
	return "current from EU Socket"
}