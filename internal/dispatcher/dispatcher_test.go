package dispatcher

import (
	"testing"
	"time"

	"github.com/sergiorra/drone-simulation/internal/models/route"
	"github.com/sergiorra/drone-simulation/internal/models/station"
	"github.com/sergiorra/drone-simulation/internal/repository/csv"
)

const (
	TestStartTime = "2011-03-22 07:47:54"
	TestFinishTime = "2011-03-22 07:48:30"
)

// TestDroneInteraction tests a drone interaction between it and the dispatcher
func TestDroneInteraction(t *testing.T) {
	repo := csv.NewRepository()
	dispatcher := NewDispatcher(repo)
	fillSampleData(dispatcher)

	go dispatcher.createDrone(dispatcher.commandsDrone1, dispatcher.reports, dispatcher.done)

	startTime, _ := time.Parse(LayoutTime, TestStartTime)
	finishTime, _ := time.Parse(LayoutTime, TestFinishTime)

	var drone1Step = 0
	dispatcher.sendCommand(6043, &drone1Step, &startTime, finishTime, &dispatcher.commandsDrone1)

	for {
		select {
		case report := <-dispatcher.reports:
			expectedIdDrone := 6043
			got := report.DroneId
			if got != expectedIdDrone {
				t.Fatalf("Expected idDrone: %v, got: %v", expectedIdDrone, got)
			}
			dispatcher.done <- true
			return
		}
	}
}

// fillSampleData fills routes and stations data with sample data
func fillSampleData(dispatcher *dispatcher) {
	dispatcher.routes = map[int][]route.Route {
		6043: {
			{
				DroneId: 6043,
				Lat: 51.478935,
				Lon: -0.172237,
				Time: time.Date(2011, time.Month(3), 22, 07, 47, 55, 0, time.UTC),
			},
		},
	}
	dispatcher.stations = []station.Station{
		{
			Name: "Acton Town",
			Lat: 51.503071,
			Lon: -0.280303,
		},
	}
}
