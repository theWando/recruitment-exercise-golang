package factory

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/theWando/car-factory/vehicle"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {

	s.adapter = &Factory{}
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSamble() {
	//code here
	// Assert
	s.Assert().Equal(1, 1)
}

func TestFactory_generateVehicleLots(t *testing.T) {
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
			fa := Factory{}
			got := fa.generateVehicleLots(tt.args.amountOfVehicles)
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
