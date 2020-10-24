package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/sergiorra/drone-simulation/internal/models/route"
	"github.com/sergiorra/drone-simulation/internal/models/station"
)

const (
	layout string = "2006-01-02 15:04:05"
)

type repository struct {
}

type StationRouteRepo interface {
	route.Repository
	station.Repository
}

// NewRepository initialize csv repository
func NewRepository() StationRouteRepo {
	return &repository{}
}

// GetRoutes fetch routes data from csv
func (r *repository) GetRoutes(dronesIds []int) (map[int][]route.Route, error) {

	wg := sync.WaitGroup{}
	wg.Add(len(dronesIds))

	routes := make(map[int][]route.Route)

	for _, val := range dronesIds {
		go readRouteFile(fmt.Sprintf("internal/data/%d.csv", val), &routes, val, &wg)
	}

	wg.Wait()

	return routes, nil
}

// readRouteFile reads csv file line by line
func readRouteFile(path string, routes *map[int][]route.Route, idDrone int, wg *sync.WaitGroup) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)

	var data []route.Route

	for values, _ := reader.Read(); values != nil; values = readLine(reader) {
		droneId, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		lat, err := strconv.ParseFloat(values[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		lon, err := strconv.ParseFloat(values[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		time, err := time.Parse(layout, values[3])
		if err != nil {
			log.Fatal(err)
		}

		route := route.NewRoute(
			droneId,
			lat,
			lon,
			time,
		)

		data = append(data, route)
	}
	(*routes)[idDrone] = data
	wg.Done()
}

// GetStations fetch stations data from csv
func (r *repository) GetStations() ([]station.Station, error) {
	f, err := os.Open("internal/data/tube.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)

	var stations []station.Station

	for values := readLine(reader); values != nil; values = readLine(reader) {
		lat, err := strconv.ParseFloat(values[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		lon, err := strconv.ParseFloat(values[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		route := station.NewStation(
			values[0],
			lat,
			lon,
		)

		stations = append(stations, route)
	}

	return stations, nil
}

func readLine(reader *csv.Reader) (line []string) {
	line, _ = reader.Read()
	return
}