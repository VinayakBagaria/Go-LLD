package main

import (
	"go-lld/parkinglot"
	"go-lld/vendingmachine"
)

const decision = "vendingMachine"

func main() {
	switch decision {
	case "parkingLot":
		parkinglot.DoWork()

	case "vendingMachine":
		vendingmachine.DoWork()
	}
}
