package moviebooking

import (
	"fmt"
	"sync/atomic"
	"time"
)

type MovieBookingService struct {
	movies       []*Movie
	theaters     []*Theater
	shows        map[string]*Show
	bookings     map[string]*Booking
	bookingCount int64
}

func GetMovieBookingService() *MovieBookingService {
	return &MovieBookingService{
		movies:       []*Movie{},
		theaters:     []*Theater{},
		shows:        map[string]*Show{},
		bookings:     map[string]*Booking{},
		bookingCount: 0,
	}
}

func (m *MovieBookingService) AddMovie(movie *Movie) {
	m.movies = append(m.movies, movie)
}

func (m *MovieBookingService) AddTheater(theater *Theater) {
	m.theaters = append(m.theaters, theater)
}

func (m *MovieBookingService) AddShow(show *Show) {
	m.shows[show.ID] = show
}

func (m *MovieBookingService) GetShow(showID string) *Show {
	return m.shows[showID]
}

func (m *MovieBookingService) BookTickets(user *User, show *Show, selectedSeats []*Seat) (*Booking, error) {
	// check seat availability
	for _, seat := range selectedSeats {
		showSeat, exists := show.Seats[seat.ID]
		if !exists || showSeat.GetStatus() != SeatStatusAvailable {
			return nil, fmt.Errorf("seat %s is not available", seat.ID)
		}
	}

	// mark seats as booked
	for _, seat := range selectedSeats {
		show.Seats[seat.ID].SetStatus(SeatStatusBooked)
	}

	// calculate price
	var totalPrice float64
	for _, seat := range selectedSeats {
		totalPrice += seat.GetPrice()
	}

	bookingID := m.generateBookingID()

	// Create booking
	booking := NewBooking(bookingID, user, show, selectedSeats, totalPrice, BookingStatusPending)
	m.bookings[bookingID] = booking

	return booking, nil
}

func (m *MovieBookingService) ConfirmBooking(bookingID string) error {
	booking, exists := m.bookings[bookingID]
	if !exists {
		return fmt.Errorf("booking not found")
	}

	if booking.GetStatus() != BookingStatusPending {
		return fmt.Errorf("booking is not in pending state")
	}

	booking.SetStatus(BookingStatusConfirmed)
	return nil
}

func (m *MovieBookingService) CancelBooking(bookingID string) error {
	booking, exists := m.bookings[bookingID]
	if !exists {
		return fmt.Errorf("booking not found")
	}

	if booking.GetStatus() == BookingStatusCancelled {
		return fmt.Errorf("booking is already cancelled")
	}

	booking.SetStatus(BookingStatusCancelled)

	for _, seat := range booking.Seats {
		booking.Show.Seats[seat.ID].SetStatus(SeatStatusAvailable)
	}

	return nil
}

func (m *MovieBookingService) generateBookingID() string {
	count := atomic.AddInt64(&m.bookingCount, 1)
	return fmt.Sprintf("BKG%s%06d", time.Now().Format("20060102150405"), count)
}
