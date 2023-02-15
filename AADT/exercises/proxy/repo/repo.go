package repo

import (
	"errors"
	"fmt"
	"time"
)

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) UpdateAllBorn(year int) (*string, error) {
	fmt.Println("called")
	if time.Now().UnixMilli()&2 == 0 {
		return nil, errors.New("Error")
	}

	success := "call passed"

	return &success, nil
}
