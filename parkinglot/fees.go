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

type VehicleRateFeeStrategy struct {
	rateCard map[VehicleType]int
}

func NewVehicleRateFeeStrategy() Fee {
	return &VehicleRateFeeStrategy{rateCard: map[VehicleType]int{
		CAR:   20,
		BIKE:  10,
		TRUCK: 30,
	}}
}

func (vehicleFee *VehicleRateFeeStrategy) CalculateFee(parkingTicket *ParkingTicket) int {
	return parkingTicket.GetElapsedTime() * vehicleFee.rateCard[parkingTicket.GetVehicle().GetType()]
}
