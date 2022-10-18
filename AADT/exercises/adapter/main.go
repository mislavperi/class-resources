package main

import (
	"example.com/adapter/adapters"
	"example.com/adapter/charger"
	"example.com/adapter/laptop"
	"example.com/adapter/sockets"
)

func main() {
	laptop := laptop.NewLaptop()
	euSocket := sockets.NewEuSocket()
	gbSocket := sockets.NewGbSocket()
	charger := charger.NewCharger()

	adapter := adapters.NewGbAdapter(gbSocket, euSocket)
	adaptedPower := adapter.AdaptToEuPower()	

	charger.Charge(laptop, adaptedPower)
}
