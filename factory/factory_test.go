package factory

import (
	"reflect"
	"testing"

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
			if got := testCar(&tt.args.car); got != tt.want {
				t.Errorf("testCar() = %v, want %v", got, tt.want)
			}
		})
	}
}
