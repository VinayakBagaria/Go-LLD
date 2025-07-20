package main

import (
	"go-lld/parkinglot"
	"go-lld/splitwise"
	"go-lld/vendingmachine"
)

const decision = "splitwise"

func main() {
	switch decision {
	case "parkingLot":
		parkinglot.DoWork()

	case "vendingMachine":
		vendingmachine.DoWork()

	case "splitwise":
		splitwise.DoWork()
	}
}
