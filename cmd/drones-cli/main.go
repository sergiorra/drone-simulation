package main

import (
	"fmt"
	"github.com/sergiorra/drone-simulation/internal/dispatcher"
	"github.com/sergiorra/drone-simulation/internal/repository/csv"
)

var DronesIds = []int {5937, 6043}

func main() {
	fmt.Println("Starting execution...")
	repo := csv.NewRepository()
	dispatcher := dispatcher.NewDispatcher(repo)
	dispatcher.GetData(DronesIds)
	dispatcher.CreateDrones()
	dispatcher.Start()
}
