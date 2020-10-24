package station

// Station representation of tube station into data struct
type Station struct {
	Name    string
	Lat     float64
	Lon  	float64
}

// NewStation initialize struct station
func NewStation(name string, lat, lon float64) (s Station) {
	s = Station{
		Name: 	name,
		Lat:    lat,
		Lon:  	lon,
	}
	return
}

// Repository definition of methods to access a data station
type Repository interface {
	GetStations() ([]Station, error)
}