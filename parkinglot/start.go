// https://github.com/ashishps1/awesome-low-level-design/blob/main/problems/parking-lot.md
package parkinglot

import (
	"fmt"
	"time"
)

func DoWork() {
	parkingLot := GetParkingLot()
	parkingLot.SetFeeStrategy(NewFlatRateFeeStrategy())
	parkingLot.AddLevel(NewLevel(1, 10))
	// parkingLot.AddLevel(NewLevel(2, 8))

	car := NewCar()
	bike := NewBike()
	truck := NewTruck()

	fmt.Printf("Ticket ID: %d\n", parkingLot.ParkVehicle(car).GetTicketId())
	fmt.Printf("Ticket ID: %d\n", parkingLot.ParkVehicle(bike).GetTicketId())
	fmt.Printf("Ticket ID: %d\n", parkingLot.ParkVehicle(truck).GetTicketId())

	// parkingLot.GetAvailability()

	time.Sleep(1 * time.Second)
	fees, _ := parkingLot.UnParkVehicle(bike)
	fmt.Println(fees)

	// parkingLot.GetAvailability()
}
