package dispatcher

import (
	"errors"
	"fmt"
	"time"

	"github.com/sergiorra/drone-simulation/internal/models/command"
	"github.com/sergiorra/drone-simulation/internal/models/report"
	"github.com/sergiorra/drone-simulation/internal/models/route"
	"github.com/sergiorra/drone-simulation/internal/models/station"
	"github.com/sergiorra/drone-simulation/internal/repository/csv"

	"github.com/sergiorra/drone-simulation/pkg/color"
	"github.com/sergiorra/drone-simulation/pkg/geolocation"
)

const (
	StartTime = "2011-03-22 07:47:54"
	FinishTime = "2011-03-22 08:10:00"
	LayoutTime = "2006-01-02 15:04:05"
	MaxDistanceToStation = 350
)

// dispatcher representation of dispatcher into data struct
type dispatcher struct {
	repo 			csv.StationRouteRepo
	routes 			map[int][]route.Route
	stations 		[]station.Station
	commandsDrone1 	chan command.Command
	commandsDrone2 	chan command.Command
	reports 		chan report.Report
	done 			chan bool
}

// NewDispatcher initialize dispatcher
func NewDispatcher(repo csv.StationRouteRepo) *dispatcher {
	return &dispatcher{
		repo: repo,
		commandsDrone1: make(chan command.Command, 10),
		commandsDrone2: make(chan command.Command, 10),
		reports: make(chan report.Report),
		done: make(chan bool),
	}
}

// GetData get necessary data about the routes and stations
func (d *dispatcher) GetData(dronesIds []int) {
	d.routes, _ = d.repo.GetRoutes(dronesIds)
	d.stations, _ = d.repo.GetStations()
}

func (d *dispatcher) CreateDrones() {
	go d.createDrone(d.commandsDrone1, d.reports, d.done)
	go d.createDrone(d.commandsDrone2, d.reports, d.done)
}

// createDrone creates a drone to send reports and receive commands by the dispatcher
func (d *dispatcher) createDrone(commands <-chan command.Command, reports chan<- report.Report, done <-chan bool) {
	for {
		select {
		case <-done:
			return
		case nextCommand := <-commands:
			printCommandReceived(nextCommand)
			if nextCommand.Route.Time.After(nextCommand.FinishTime) {
				break
			}
			if nextCommand.CountdownTime > 0 {
				time.Sleep(time.Duration(nextCommand.CountdownTime) * time.Second)
			}

			lastReport := report.NewReport(nextCommand.Route.DroneId, nextCommand.Route.Time, 30)

			station, err := d.findStation(nextCommand.Route.Lat, nextCommand.Route.Lon)
			if err != nil {
				lastReport.StationFound = false
			} else {
				lastReport.StationName = station.Name
			}

			reports <- lastReport
		}
	}
}

// Start starts the dispatcher execution to send commands to the drones and receive reports by the drones
func (d *dispatcher) Start() {
	startTime, _ := time.Parse(LayoutTime, StartTime)
	finishTime, _ := time.Parse(LayoutTime, FinishTime)

	var drone1Step, drone2Step = 0, 0
	d.sendCommand(5937, &drone1Step, &startTime, finishTime, &d.commandsDrone1)
	d.sendCommand(6043, &drone2Step, &startTime, finishTime, &d.commandsDrone2)

	ticker := time.NewTicker(time.Second)

	for  {
		select {
		case <-ticker.C:
			startTime = startTime.Add(time.Second)
			if startTime.After(finishTime) {
				fmt.Println("Time Finished!!")
				d.done <- true
				ticker.Stop()
				return
			}
		case report := <-d.reports:
			if report.StationFound {
				reportTraffic(report)
			}
			switch report.DroneId {
			case 5937:
				drone1Step++
				d.sendCommand(5937, &drone1Step, &startTime, finishTime, &d.commandsDrone1)
			case 6043:
				drone2Step++
				d.sendCommand(6043, &drone2Step, &startTime, finishTime, &d.commandsDrone2)
			}
		}
	}
}

// findStation search for a nearby station
func (d *dispatcher) findStation(lat float64, lon float64) (station.Station, error) {
	for _, station  := range d.stations {
		if geolocation.CheckDistance(lat, lon, station.Lat, station.Lon) <= MaxDistanceToStation {
			return station, nil
		}
	}

	return station.Station{}, errors.New("no stations found")
}

// sendCommand sends a new command with the new route and the countdown to reach it to a drone
func (d *dispatcher) sendCommand(droneId int, droneStep *int, startTime *time.Time, finishTime time.Time, commandsDrone *chan command.Command) {
	nextRoute := d.routes[droneId][*droneStep]
	countdown := int(nextRoute.Time.Sub(*startTime).Seconds())
	nextCommand := command.NewCommand(nextRoute, countdown, finishTime)
	*commandsDrone <- nextCommand
}

// printCommandReceived prints a command received by a drone
func printCommandReceived(command command.Command) {
	fmt.Printf("Command received by droneId %v, going to coordinates %v, %v at time %v\n", command.Route.DroneId,
				command.Route.Lat, command.Route.Lon, command.Route.Time )
}

// reportTraffic prints data about traffic report
func reportTraffic(report report.Report) {
	fmt.Println(color.Yellow + "-------------------------" + color.Reset)
	fmt.Println(color.Yellow + "TRAFFIC REPORT:" + color.Reset)
	fmt.Printf("   Drone Id: %v\n", report.DroneId)
	fmt.Printf("   Station Name: %v\n", report.StationName)
	fmt.Printf("   Time: %v\n", report.Time)
	fmt.Printf("   Speed: %v\n", report.Speed)
	fmt.Printf("   Traffic Condition: %v\n", report.Condition)
	fmt.Println(color.Yellow + "-------------------------" + color.Reset)
}