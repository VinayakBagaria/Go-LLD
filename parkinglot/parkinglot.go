package parkinglot

import "fmt"

type ParkingLot struct {
	levels []*Level
}

var instance *ParkingLot

func GetParkingLot() *ParkingLot {
	if instance == nil {
		instance = &ParkingLot{levels: []*Level{}}
	}
	return instance
}

func (p *ParkingLot) AddLevel(level *Level) {
	p.levels = append(p.levels, level)
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.levels {
		if level.ParkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) UnParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.levels {
		if level.UnParkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) GetAvailability() {
	for _, level := range p.levels {
		level.GetAvailability()
	}
	fmt.Println()
}
