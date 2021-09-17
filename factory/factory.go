package factory

import (
	"github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"

	"github.com/theWando/car-factory/assemblyspot"
	"github.com/theWando/car-factory/vehicle"
)

type Factory struct {
	pool *ants.Pool
}

func New(assemblySpots int) (*Factory, error) {
	pool, err := ants.NewPool(assemblySpots)
	if err != nil {
		return nil, err
	}
	factory := &Factory{
		pool: pool,
	}

	return factory, nil
}

func (f Factory) StartAssemblingProcess(amountOfVehicles int, out chan vehicle.Car) {
	defer close(out)
	defer f.pool.Release()
	vehicleList := generateVehicleLots(amountOfVehicles)
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

			car.TestingLog = testCar(&car)
			car.AssembleLog = idleSpot.GetAssembledLogs()
			out <- car
		})
		if err != nil {
			log.Error(err)
			break
		}
	}
}

func generateVehicleLots(amountOfVehicles int) []vehicle.Car {
	vehicles := make([]vehicle.Car, amountOfVehicles)

	for i := 0; i < amountOfVehicles; i++ {
		vehicles[i] = vehicle.Car{
			Id:            i,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
		}
	}

	return vehicles
}

func testCar(car *vehicle.Car) string {
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
