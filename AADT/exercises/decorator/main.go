package main

import (
	"example.com/decorator/checkup"
	"example.com/decorator/carwash"
)

func main() {
	checkup := checkup.NewCheckUp()
	checkup.DoCheckup()
	newCarWash := carwash.NewCarWash(checkup)
	newCarWash.WashCar()
}
