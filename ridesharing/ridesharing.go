package ridesharing

import (
	"fmt"
	"time"
)

type RideService struct {
	passengers map[int]*Passenger
	drivers    map[int]*Driver
	rides      map[int]*Ride
}

func NewRideService() *RideService {
	return &RideService{
		passengers: make(map[int]*Passenger),
		drivers:    make(map[int]*Driver),
		rides:      make(map[int]*Ride),
	}
}

func (rs *RideService) AddPassenger(passenger *Passenger) {
	rs.passengers[passenger.ID] = passenger
}

func (rs *RideService) AddDriver(driver *Driver) {
	rs.drivers[driver.ID] = driver
}

func (rs *RideService) RequestRide(passenger *Passenger, source, destination *Location) *Ride {
	ride := &Ride{
		ID:          rs.generateRideId(),
		Passenger:   passenger,
		Source:      source,
		Destination: destination,
		Status:      Requested,
	}
	rs.notifyDrivers(ride)
	return ride
}

func (rs *RideService) AcceptRide(driver *Driver, ride *Ride) {
	if ride.Status == Requested {
		ride.Driver = driver
		ride.Status = Accepted
		driver.Status = Busy
		rs.notifyPassenger(ride)
	}
}

func (rs *RideService) StartRide(ride *Ride) {
	if ride.Status == Accepted {
		ride.Status = InProgress
		rs.notifyPassenger(ride)
	}
}

func (rs *RideService) CompleteRide(ride *Ride) {
	if ride.Status == InProgress {
		ride.Status = Completed
		ride.Driver.Status = Available
		ride.Fare = rs.calculateFare(ride)
		rs.notifyDriver(ride)
		rs.notifyPassenger(ride)
	}
}

func (rs *RideService) CancelRide(ride *Ride) {
	if ride.Status == Requested || ride.Status == Accepted {
		ride.Status = Cancelled
		rs.notifyDriver(ride)
		rs.notifyPassenger(ride)
		if ride.Driver != nil {
			ride.Driver.Status = Available
			ride.Driver = nil
		}
	}
}

func (rs *RideService) notifyDrivers(ride *Ride) {
	for _, driver := range rs.drivers {
		if driver.Status == Available && driver.Location.CalculateDistance(ride.Source) <= 5.0 {
			fmt.Printf("Notifying driver %s about ride request: %d\n", driver.Name, ride.ID)
		}
	}
}

func (rs *RideService) notifyDriver(ride *Ride) {
	if ride.Driver != nil {
		message := ""
		switch ride.Status {
		case Completed:
			message = fmt.Sprintf("Ride completed. Fare: $%.2f", ride.Fare)
		case Cancelled:
			message = "Ride cancelled by passenger"
		}
		fmt.Printf("Notifying driver %s: %s\n", ride.Driver.Name, message)
	}
}

func (rs *RideService) notifyPassenger(ride *Ride) {
	message := ""
	switch ride.Status {
	case Accepted:
		message = fmt.Sprintf("Your ride has been accepted by driver: %s", ride.Driver.Name)
	case InProgress:
		message = "Your ride is in progress"
	case Completed:
		message = fmt.Sprintf("Your ride has been completed. Fare: $%.2f", ride.Fare)
	case Cancelled:
		message = "Your ride has been cancelled"
	}
	fmt.Printf("Notifying passenger %s: %s\n", ride.Passenger.Name, message)
}

func (rs *RideService) calculateFare(ride *Ride) float64 {
	baseFare := 2.0
	perKmFare := 1.5
	distance := ride.Source.CalculateDistance(ride.Destination)
	return baseFare + (distance * perKmFare)
}

func (rs *RideService) generateRideId() int {
	return int(time.Now().Unix())
}
