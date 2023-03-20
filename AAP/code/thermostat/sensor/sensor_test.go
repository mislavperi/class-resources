package sensor_test

import (
	"testing"

	"example.com/thermostat/sensor"
	"example.com/thermostat/sensor/mocks"
	"example.com/thermostat/thermo"
)

func SensorTest(t *testing.T) {
	singalMock := &mocks.SignalL{}
	sns := sensor.NewSensor(
		10.5, thermo.NewThermo(9),
		singalMock,
	)
	sns.CheckThreshhold(sns.GetCurrentTemperature())
}
