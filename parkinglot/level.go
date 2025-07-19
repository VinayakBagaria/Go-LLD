package parkinglot

import "fmt"

type Level struct {
	floor        int
	parkingSpots []*ParkingSpot
}

func NewLevel(floor, numSpots int) *Level {
	level := &Level{floor: floor}

	bikeSpots := int(float64(numSpots) * 0.5)
	carSpots := int(float64(numSpots) * 0.4)

	for i := 1; i <= bikeSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(i, BIKE))
	}

	for i := bikeSpots + 1; i <= bikeSpots+carSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(i, CAR))
	}

	for i := bikeSpots + carSpots + 1; i <= numSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(i, TRUCK))
	}

	return level
}

func (level *Level) ParkVehicle(vehicle Vehicle) *ParkingTicket {
	for _, spot := range level.parkingSpots {
		if spot.ParkVehicle(vehicle) {
			parkingTicket := NewParkingTicket(vehicle, spot)
			return parkingTicket
		}
	}

	return nil
}

func (level *Level) UnParkVehicle(vehicle Vehicle) bool {
	for _, spot := range level.parkingSpots {
		if !spot.IsAvailable() && spot.GetParkedVehicle() == vehicle {
			spot.UnParkVehicle()
			return true
		}
	}

	return false
}

func (level *Level) GetAvailability() {
	occupied := 0
	for _, spot := range level.parkingSpots {
		status := "Available"
		if !spot.IsAvailable() {
			status = "Occupied"
			occupied++
		}
		fmt.Printf("Level: %d, Spot: %d, Status: %s, Type: %d\n", level.floor, spot.GetIdentifier(), status, spot.GetVehicleType())
	}

	fmt.Printf("Total Spots: %d, Occupied: %d\n", len(level.parkingSpots), occupied)
}
