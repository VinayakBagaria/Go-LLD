// https://github.com/ashishps1/awesome-low-level-design/blob/main/problems/ride-sharing-service.md
package ridesharing

import (
	"fmt"
	"time"
)

func DoWork() {
	rideService := NewRideService()

	passenger1 := &Passenger{ID: 1, Name: "John Doe", Location: &Location{Latitude: 37.7749, Longitude: -122.4194}}
	passenger2 := &Passenger{ID: 2, Name: "Jane Smith", Location: &Location{Latitude: 37.7860, Longitude: -122.4070}}
	rideService.AddPassenger(passenger1)
	rideService.AddPassenger(passenger2)

	driver1 := &Driver{ID: 1, Name: "Alice Johnson", LicensePlate: "ABC123", Location: &Location{Latitude: 37.7749, Longitude: -122.4194}, Status: Available}
	driver2 := &Driver{ID: 2, Name: "Bob Williams", LicensePlate: "XYZ789", Location: &Location{Latitude: 37.7860, Longitude: -122.4070}, Status: Available}
	rideService.AddDriver(driver1)
	rideService.AddDriver(driver2)

	// Passenger 1 requests a ride
	ride1 := rideService.RequestRide(passenger1, passenger1.Location, &Location{Latitude: 37.7887, Longitude: -122.4098})

	// Driver 1 accepts the ride
	rideService.AcceptRide(driver1, ride1)

	// Start and complete the ride
	rideService.StartRide(ride1)
	rideService.CompleteRide(ride1)

	time.Sleep(1 * time.Second)
	fmt.Println()

	// Passenger 2 requests a ride and cancels it
	ride2 := rideService.RequestRide(passenger2, passenger2.Location, &Location{Latitude: 37.7749, Longitude: -122.4194})
	rideService.AcceptRide(driver2, ride2)
	rideService.CancelRide(ride2)
}
