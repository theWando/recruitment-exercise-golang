package factory

import (
	"github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"

	"github.com/theWando/car-factory/assemblyspot"
	"github.com/theWando/car-factory/vehicle"
)

type Factory struct {
	AssemblingSpots chan *assemblyspot.AssemblySpot
	pool            *ants.Pool
}

func New(assemblySpots int) (*Factory, error) {
	pool, err := ants.NewPool(assemblySpots)
	if err != nil {
		return nil, err
	}
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
		pool:            pool,
	}

	totalAssemblySpots := 0

	for {
		factory.AssemblingSpots <- &assemblyspot.AssemblySpot{}

		totalAssemblySpots++

		if totalAssemblySpots >= assemblySpots {
			break
		}
	}

	return factory, nil
}

func (f Factory) Assemble(amountOfVehicles int, out chan vehicle.Car) {
	defer close(out)
	vehicleList := f.generateVehicleLots(amountOfVehicles)
	for _, parts := range vehicleList {
		parts := parts
		err := f.pool.Submit(func() {
			var idleSpot assemblyspot.AssemblySpot
			log.WithField("ID", parts.Id).Info("Assembling vehicle...")
			idleSpot.SetVehicle(&parts)
			car, err := idleSpot.AssembleVehicle()

			if err != nil {
				log.WithField("err", err).Error("failed to assemble vehicle")
			}

			car.TestingLog = f.testCar(&car)
			car.AssembleLog = idleSpot.GetAssembledLogs()
			out <- car
		})
		if err != nil {
			log.Error(err)
			break
		}
	}
}

func (Factory) generateVehicleLots(amountOfVehicles int) []vehicle.Car {
	var vehicles []vehicle.Car
	var index = 0

	for {
		vehicles = append(vehicles, vehicle.Car{
			Id:            index,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
		})

		index++

		if index >= amountOfVehicles {
			break
		}
	}

	return vehicles
}

func (f *Factory) testCar(car *vehicle.Car) string {
	logs := ""

	trace, err := car.StartEngine()
	if err == nil {
		logs += trace + ", "
	} else {
		logs += err.Error() + ", "
	}

	trace, err = car.MoveForwards(10)
	if err == nil {
		logs += trace + ", "
	} else {
		logs += err.Error() + ", "
	}

	trace, err = car.MoveForwards(10)
	if err == nil {
		logs += trace + ", "
	} else {
		logs += err.Error() + ", "
	}

	trace, err = car.TurnLeft()
	if err == nil {
		logs += trace + ", "
	} else {
		logs += err.Error() + ", "
	}

	trace, err = car.TurnRight()
	if err == nil {
		logs += trace + ", "
	} else {
		logs += err.Error() + ", "
	}

	trace, err = car.StopEngine()
	if err == nil {
		logs += trace + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
