package parkinglot

type Fee interface {
	CalculateFee(parkingTicket *ParkingTicket) int
}

type FlatRateFeeStrategy struct {
	ratePerSecond int
}

func NewFlatRateFeeStrategy() Fee {
	return &FlatRateFeeStrategy{ratePerSecond: 10}
}

func (flatFee *FlatRateFeeStrategy) CalculateFee(parkingTicket *ParkingTicket) int {
	return parkingTicket.GetElapsedTime() * flatFee.ratePerSecond
}
