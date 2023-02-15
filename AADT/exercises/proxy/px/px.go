package px

import (
	"time"
)

type ProxyImp struct {
}

type Repo interface {
	UpdateAllBorn(year int) (*string, error)
}

type fncall func(year int) (*string, error)

func NewPx() *ProxyImp {
	return &ProxyImp{}
}

func (p *ProxyImp) Backoff(retry int, fn fncall, year int) (*string, error) {
	var err error
	for i := 1; i <= retry; i++ {
		msg, err := fn(year)
		if err != nil {
			return msg, nil
		}
		time.Sleep(3 * time.Second)
	}
	return nil, err
}
