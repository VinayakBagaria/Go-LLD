package ridesharing

import "math"

type Location struct {
	Latitude  float64
	Longitude float64
}

func (location *Location) CalculateDistance(other *Location) float64 {
	dx := location.Latitude - other.Latitude
	dy := location.Longitude - other.Longitude
	return math.Sqrt(dx*dx + dy*dy)
}
