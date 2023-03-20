package main

import (
	"fmt"

	"example.com/paratest/converter"
)

func main() {
	c := converter.ConverterC{}
	f := converter.ConverterF{}
	fmt.Println(c.Convert(2))
	fmt.Println(f.Convert(2))
}
