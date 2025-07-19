package parkinglot

import (
	"errors"
	"fmt"
)

type ParkingLot struct {
	levels        []*Level
	activeTickets map[Vehicle]*ParkingTicket
	feeStrategy   Fee
}

var instance *ParkingLot

func GetParkingLot() *ParkingLot {
	if instance == nil {
		instance = &ParkingLot{levels: []*Level{}, activeTickets: make(map[Vehicle]*ParkingTicket)}
	}
	return instance
}

func (p *ParkingLot) SetFeeStrategy(feeStrategy Fee) {
	p.feeStrategy = feeStrategy
}

func (p *ParkingLot) AddLevel(level *Level) {
	p.levels = append(p.levels, level)
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) *ParkingTicket {
	for _, level := range p.levels {
		parkingTicket := level.ParkVehicle(vehicle)
		if parkingTicket != nil {
			p.activeTickets[vehicle] = parkingTicket
			return parkingTicket
		}
	}
	return nil
}

func (p *ParkingLot) UnParkVehicle(vehicle Vehicle) (int, error) {
	ticket, ok := p.activeTickets[vehicle]
	if !ok {
		return -1, errors.New("Vehicle not parked anywhere")
	}

	for _, level := range p.levels {
		if level.UnParkVehicle(vehicle) {
			ticket.MarkExit()
			delete(p.activeTickets, vehicle)
			return p.feeStrategy.CalculateFee(ticket), nil
		}
	}

	return -1, errors.New("Vehicle not parked anywhere")
}

func (p *ParkingLot) GetAvailability() {
	for _, level := range p.levels {
		level.GetAvailability()
	}
	fmt.Println()
}
