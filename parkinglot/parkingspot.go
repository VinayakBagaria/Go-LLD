package parkinglot

type ParkingSpot struct {
	identifier    int
	vehicleType   VehicleType
	parkedVehicle Vehicle
}

func NewParkingSpot(identifier int, vehicleType VehicleType) *ParkingSpot {
	return &ParkingSpot{identifier: identifier, vehicleType: vehicleType}
}

func (ps *ParkingSpot) IsAvailable() bool {
	return ps.parkedVehicle == nil
}

func (ps *ParkingSpot) ParkVehicle(vehicle Vehicle) bool {
	if ps.IsAvailable() && ps.vehicleType == vehicle.GetType() {
		ps.parkedVehicle = vehicle
		return true
	}

	return false
}

func (ps *ParkingSpot) UnParkVehicle() {
	ps.parkedVehicle = nil
}

func (ps *ParkingSpot) GetParkedVehicle() Vehicle {
	return ps.parkedVehicle
}

func (ps *ParkingSpot) GetIdentifier() int {
	return ps.identifier
}

func (ps *ParkingSpot) GetVehicleType() VehicleType {
	return ps.vehicleType
}
