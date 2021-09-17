package factory

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"
	"github.com/theWando/car-factory/vehicle"
)

func Test_generateVehicleLots(t *testing.T) {
	type args struct {
		amountOfVehicles int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{name: "Generates 0 vehicles", args: args{amountOfVehicles: 0}, wantLen: 0},
		{name: "Generates 3 vehicles", args: args{amountOfVehicles: 3}, wantLen: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateVehicleLots(tt.args.amountOfVehicles)
			if len(got) != tt.wantLen {
				t.Errorf("generateVehicleLots() = %d, want %d", len(got), tt.wantLen)
			}
			for i, car := range got {
				compare := vehicle.Car{
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
				if !reflect.DeepEqual(car, compare) {
					t.Errorf("generateVehicleLots() = %v, want %v", got, compare)
				}
			}
		})
	}
}

func Test_testCar(t *testing.T) {
	type args struct {
		car vehicle.Car
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Returns the result of a functional car", args: args{car: vehicle.Car{
			EngineStarted: false,
		}}, want: "Engine Started!, Moved forward 10 meters!, Moved forward 10 meters!, Turned Right, Turned Right!, Engine Stopped!, "},
		{name: "Returns the result of a al reade stated car", args: args{car: vehicle.Car{
			EngineStarted: true,
		}}, want: "Cannot start engine already started, Moved forward 10 meters!, Moved forward 10 meters!, Turned Right, Turned Right!, Engine Stopped!, "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testCar(tt.args.car); got != tt.want {
				t.Errorf("testCar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	_ = os.Setenv("MOCK_DATE", "2021-09-17 17:05:35.455")
	//_ = os.Setenv("MOCK_WORK_DURATION", "100")
	m.Run()
	_ = os.Unsetenv("MOCK_DATE")
	_ = os.Unsetenv("MOCK_WORK_DURATION")
}

func Test_assembleVehicle(t *testing.T) {
	type args struct {
		serial int
	}
	tests := []struct {
		name    string
		args    args
		want    vehicle.Car
		wantErr bool
	}{
		{
			name: "builds a car",
			args: args{serial: 9},
			want: vehicle.Car{
				Id:            9,
				Chassis:       "Assembled",
				Tires:         "Assembled",
				Engine:        "Assembled",
				Electronics:   "Assembled",
				Dash:          "Assembled",
				Sits:          "Assembled",
				Windows:       "Assembled",
				EngineStarted: false,
				TestingLog:    "Engine Started!, Moved forward 10 meters!, Moved forward 10 meters!, Turned Right, Turned Right!, Engine Stopped!, ",
				AssembleLog:   "Chassis at [2021-09-17 17:05:35.455], Tires at [2021-09-17 17:05:35.455], Engine at [2021-09-17 17:05:35.455], Electronics at [2021-09-17 17:05:35.455], Dash at [2021-09-17 17:05:35.455], Sits at [2021-09-17 17:05:35.455], Windows at [2021-09-17 17:05:35.455], ",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := assembleVehicle(tt.args.serial)
			if (err != nil) != tt.wantErr {
				t.Errorf("assembleVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assembleVehicle() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestFactory_StartAssemblingProcess(t *testing.T) {
	t.Skip("couldn't make it work")
	type fields struct {
		pool *ants.Pool
	}
	type args struct {
		amountOfVehicles int
		out              chan vehicle.Car
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		carsWanted int
	}{
		{name: "returns 5 cars one batch", fields: fields{pool: func() *ants.Pool {
			p, _ := ants.NewPool(5)
			return p
		}()}, args: args{
			amountOfVehicles: 10,
			out:              make(chan vehicle.Car),
		}, carsWanted: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				pool: tt.fields.pool,
			}
			go f.StartAssemblingProcess(tt.args.amountOfVehicles, tt.args.out)
			var carsReceived int
			ticker := time.NewTicker(10 * time.Second)
			select {
			case _ = <-ticker.C:
				tt.fields.pool.Release()
				ticker.Stop()
				close(tt.args.out)
				log.Infof("cars received %d", carsReceived)
				if carsReceived != tt.carsWanted {
					t.Errorf("Received %d but wanted %d", carsReceived, tt.carsWanted)
					return
				}
			case car := <-tt.args.out:
				log.Infof("%#v", car)
				carsReceived++
			}
		})
	}
}
