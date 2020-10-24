package report

import (
	"math/rand"
	"time"
)

// Report representation of report into data struct
type Report struct {
	DroneId    		int
	Time     		time.Time
	Speed  			float64
	Condition 		ConditionType
	StationFound	bool
	StationName 	string
}

// NewReport initialize struct report
func NewReport(droneId int, time time.Time, speed float64) (r Report) {
	r = Report{
		DroneId: 		droneId,
		Time:      		time,
		Speed:  		speed,
		Condition:  	NewConditionType(),
		StationFound: 	true,
	}
	return
}

type ConditionType int

const (
	Heavy ConditionType = iota
	Light
	Moderate
)

func (t ConditionType) String() string {
	return toString[t]
}

// NewConditionType initialize a random type from enum conditionTypes
func NewConditionType() ConditionType {
	random := ConditionType(rand.Intn(3))
	conditionType := toID[random.String()]
	return conditionType
}

var toString = map[ConditionType]string{
	Heavy:         "Heavy",
	Light:         "Light",
	Moderate:      "Moderate",
}

var toID = map[string]ConditionType{
	"Heavy":        Heavy,
	"Light":        Light,
	"Moderate":  	Moderate,
}