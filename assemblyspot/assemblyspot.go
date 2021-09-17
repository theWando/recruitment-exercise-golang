package assemblyspot

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/theWando/car-factory/vehicle"
)

const (
	layout           = "2006-01-02 15:04:05.000"
	assembleDuration = 1 * time.Second
)

type AssemblySpot struct {
	vehicleToAssemble *vehicle.Car
	assemblyLog       string
}

func (s *AssemblySpot) SetVehicle(v vehicle.Car) {
	s.vehicleToAssemble = &v
}

func (s *AssemblySpot) GetAssembledVehicle() vehicle.Car {
	return *s.vehicleToAssemble
}

func (s *AssemblySpot) GetAssembledLogs() string {
	return s.assemblyLog
}

//hint: improve this function to execute this process concurrenlty
func (s *AssemblySpot) AssembleVehicle() error {
	if s.vehicleToAssemble == nil {
		return errors.New("no vehicle set to start assembling")
	}

	s.assembleChassis()
	s.assembleTires()
	s.assembleEngine()
	s.assembleElectronics()
	s.assembleDash()
	s.assembleSeats()
	s.assembleWindows()

	return nil
}

func (s *AssemblySpot) assembleChassis() {
	s.vehicleToAssemble.Chassis = "Assembled"
	doSomeWork()
	s.assemblyLog += fmt.Sprintf("Chassis at [%s], ", getNow())
}

func doSomeWork() {
	if mockWork := os.Getenv("MOCK_WORK_DURATION"); mockWork != "" {
		value, err := strconv.Atoi(mockWork)
		if err == nil {
			duration := time.Duration(value) * time.Millisecond
			time.Sleep(duration)
			return
		}
		log.Errorf("failed to convert value, using default")
	}
	time.Sleep(assembleDuration)
}

func (s *AssemblySpot) assembleTires() {
	s.vehicleToAssemble.Tires = "Assembled"
	doSomeWork()
	s.assemblyLog += fmt.Sprintf("Tires at [%s], ", getNow())
}

func (s *AssemblySpot) assembleEngine() {
	s.vehicleToAssemble.Engine = "Assembled"
	doSomeWork()
	s.assemblyLog += fmt.Sprintf("Engine at [%s], ", getNow())
}

func (s *AssemblySpot) assembleElectronics() {
	s.vehicleToAssemble.Electronics = "Assembled"
	doSomeWork()
	s.assemblyLog += fmt.Sprintf("Electronics at [%s], ", getNow())
}

func (s *AssemblySpot) assembleDash() {
	s.vehicleToAssemble.Dash = "Assembled"
	doSomeWork()
	s.assemblyLog += fmt.Sprintf("Dash at [%s], ", getNow())
}

func (s *AssemblySpot) assembleSeats() {
	s.vehicleToAssemble.Sits = "Assembled"
	doSomeWork()
	s.assemblyLog += fmt.Sprintf("Sits at [%s], ", getNow())
}

func (s *AssemblySpot) assembleWindows() {
	s.vehicleToAssemble.Windows = "Assembled"
	doSomeWork()
	s.assemblyLog += fmt.Sprintf("Windows at [%s], ", getNow())
}

func getNow() string {
	if mockDate := os.Getenv("MOCK_DATE"); mockDate != "" {
		parse, err := time.Parse(layout, mockDate)
		if err != nil {
			return "Failed to parse mock date"
		}
		return parse.Format(layout)
	}
	return time.Now().Format(layout)
}
