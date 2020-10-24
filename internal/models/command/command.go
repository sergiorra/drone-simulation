package command

import (
	"time"

	"github.com/sergiorra/drone-simulation/internal/models/route"
)

// Command representation of command into data struct
type Command struct {
	Route 			route.Route
	CountdownTime  	int
	FinishTime 		time.Time
}

// NewCommand initialize struct command
func NewCommand(route route.Route, countdownTime int, finishTime time.Time) (c Command) {
	c = Command{
		Route: 			route,
		CountdownTime:	countdownTime,
		FinishTime: 	finishTime,
	}
	return
}
