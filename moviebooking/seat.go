package moviebooking

import (
	"fmt"
	"sync"
)

type Seat struct {
	ID     string
	Row    int
	Column int
	Type   SeatType
	Price  float64
	status SeatStatus
	mu     sync.RWMutex
}

func NewSeat(id string, row, column int, seatType SeatType, price float64, status SeatStatus) *Seat {
	return &Seat{
		ID:     id,
		Row:    row,
		Column: column,
		Type:   seatType,
		Price:  price,
		status: status,
	}
}

func (s *Seat) GetStatus() SeatStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.status
}

func (s *Seat) SetStatus(status SeatStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status = status
}

func (s *Seat) GetPrice() float64 {
	return s.Price
}

func CreateSeats(rows, columns int) map[string]*Seat {
	seats := make(map[string]*Seat)
	for row := 1; row <= rows; row++ {
		for col := 1; col <= columns; col++ {
			seatID := fmt.Sprintf("%d-%d", row, col)
			seatType := SeatTypeNormal
			price := 100.0

			if row <= 2 {
				seatType = SeatTypePremium
				price = 150.0
			}

			seats[seatID] = NewSeat(
				seatID,
				row,
				col,
				seatType,
				price,
				SeatStatusAvailable,
			)
		}
	}
	return seats
}
