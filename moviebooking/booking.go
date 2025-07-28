package moviebooking

type Booking struct {
	ID         string
	User       *User
	Show       *Show
	Seats      []*Seat
	TotalPrice float64
	Status     BookingStatus
}

func NewBooking(id string, user *User, show *Show, seats []*Seat, totalPrice float64, status BookingStatus) *Booking {
	return &Booking{
		ID:         id,
		User:       user,
		Show:       show,
		Seats:      seats,
		TotalPrice: totalPrice,
		Status:     status,
	}
}

func (b *Booking) GetStatus() BookingStatus {
	return b.Status
}

func (b *Booking) SetStatus(status BookingStatus) {
	b.Status = status
}
