package parkinglot

import (
	"math/rand"
	"time"
)

type ParkingTicket struct {
	ticketId    int
	vehicle     Vehicle
	parkingSpot *ParkingSpot
	entry       time.Time
	exit        time.Time
}

func NewParkingTicket(vehicle Vehicle, parkingSpot *ParkingSpot) *ParkingTicket {
	return &ParkingTicket{ticketId: rand.Int(), vehicle: vehicle, parkingSpot: parkingSpot, entry: time.Now()}
}

func (ticket *ParkingTicket) MarkExit() {
	ticket.exit = time.Now()
}

func (ticket *ParkingTicket) GetElapsedTime() int {
	return int(ticket.exit.Sub(ticket.entry).Seconds())
}

func (ticket *ParkingTicket) GetTicketId() int {
	return ticket.ticketId
}
