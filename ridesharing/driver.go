package ridesharing

type Driver struct {
	ID           int
	Name         string
	LicensePlate string
	Location     *Location
	Status       DriverStatus
}
