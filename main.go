package main

import (
	"go-lld/parkinglot"
	"go-lld/ridesharing"
	"go-lld/splitwise"
	"go-lld/vendingmachine"
)

const decision = "rideSharing"

func main() {
	switch decision {
	case "parkingLot":
		parkinglot.DoWork()

	case "vendingMachine":
		vendingmachine.DoWork()

	case "splitwise":
		splitwise.DoWork()

	case "rideSharing":
		ridesharing.DoWork()
	}
}
