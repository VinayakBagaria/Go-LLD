// https://github.com/ashishps1/awesome-low-level-design/blob/main/problems/parking-lot.md
package parkinglot

func DoWork() {
	parkingLot := GetParkingLot()
	parkingLot.AddLevel(NewLevel(1, 10))
	// parkingLot.AddLevel(NewLevel(2, 8))

	car := NewCar()
	bike := NewBike()
	truck := NewTruck()

	parkingLot.ParkVehicle(car)
	parkingLot.ParkVehicle(bike)
	parkingLot.ParkVehicle(truck)

	parkingLot.GetAvailability()

	parkingLot.UnParkVehicle(bike)

	parkingLot.GetAvailability()
}
