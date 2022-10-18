package sockets

type GbSocket struct {
}

func NewGbSocket() *GbSocket {
	return &GbSocket{}
}

func (g *GbSocket) GiveGbCurrent() string {
	return "current from GB Socket"
}