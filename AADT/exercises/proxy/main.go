package main

import (
	"fmt"

	"example.com/proxy/px"
	"example.com/proxy/repo"
)

type Repo interface {
	UpdateAllBorn(year int) (*string, error)
}

func main() {
	r := repo.NewRepo()
	px := px.NewPx()
	msg, err := px.Backoff(10, r.UpdateAllBorn, 2000)
	fmt.Println(msg, err)
}
