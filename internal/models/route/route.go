package route

import (
	"time"
)

// Route representation of route into data struct
type Route struct {
	DroneId    	int
	Lat     	float64
	Lon  		float64
	Time  		time.Time
}

// NewRoute initialize struct route
func NewRoute(droneId int, lat, lon float64, time time.Time) (r Route) {
	r = Route{
		DroneId: 	droneId,
		Lat:      	lat,
		Lon:  		lon,
		Time:      	time,
	}
	return
}

// Repository definition of methods to access a data route
type Repository interface {
	GetRoutes([]int) (map[int][]Route, error)
}