package main

import (
	"go-lld/parkinglot"
)

const decision = "parkingLot"

func main() {
	switch decision {
	case "parkingLot":
		parkinglot.DoWork()
	}
}
