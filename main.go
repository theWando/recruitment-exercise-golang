package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/theWando/car-factory/factory"
	"github.com/theWando/car-factory/vehicle"
)

const (
	carsAmount       = 100
	assemblyCapacity = 5
)

func main() {
	log.SetOutput(os.Stdout)

	app, err := factory.New(assemblyCapacity)
	if err != nil {
		log.Error(err)
	}

	//Hint: change appropriately for making factory give each vehicle once assembled,
	//even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id

	carChannel := make(chan vehicle.Car)
	go app.Assemble(carsAmount, carChannel)
	for car := range carChannel {
		log.WithFields(map[string]interface{}{
			"TestingLog":  car.TestingLog,
			"AssembleLog": car.AssembleLog,
		}).Infof("car ID %d ready", car.Id)
	}
}
