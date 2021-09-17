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
	pool, err := ants.NewPool(assemblySpots, func(opts *ants.Options) {
		opts.Logger = log.StandardLogger()
		opts.PanicHandler = func(i interface{}) {
			log.WithField("error", i).Errorf("received panic from factory")
		}
	})
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
	for i := 0; i < amountOfVehicles; i++ {
		serial := i
		err := f.pool.Submit(func() {
			car, _ := assembleVehicle(serial)
			out <- car
		})
		if err != nil {
			return
		}
	}
}

func assembleVehicle(serial int) (vehicle.Car, error) {
	const notSet = "NotSet"
	var spot assemblyspot.AssemblySpot
	log.WithField("ID", serial).Info("Assembling vehicle...")
	spot.SetVehicle(vehicle.Car{
		Id:            serial,
		Chassis:       notSet,
		Tires:         notSet,
		Engine:        notSet,
		Electronics:   notSet,
		Dash:          notSet,
		Sits:          notSet,
		Windows:       notSet,
		EngineStarted: false,
	})

	err := spot.AssembleVehicle()
	if err != nil {
		log.WithField("err", err).Error("failed to assemble vehicle")
		return vehicle.Car{}, err
	}

	car := spot.GetAssembledVehicle()
	car.TestingLog = testCar(car)
	car.AssembleLog = spot.GetAssembledLogs()
	return car, nil
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

func testCar(car vehicle.Car) string {
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
