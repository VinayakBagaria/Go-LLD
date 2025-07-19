package parkinglot

type VehicleType int

const (
	CAR VehicleType = iota
	BIKE
	TRUCK
)

type Vehicle interface {
	GetType() VehicleType
}

type BaseVehicle struct {
	vehicleType VehicleType
}

func (v BaseVehicle) GetType() VehicleType {
	return v.vehicleType
}
