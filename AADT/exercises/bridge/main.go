package main

import (
	"example.com/bridge/shape"
	"example.com/bridge/output"
)

func main() {
	shape := shape.NewShape("Triangle")
	output := output.NewOutput("Monitor")

	output.DrawShape(shape)
}
