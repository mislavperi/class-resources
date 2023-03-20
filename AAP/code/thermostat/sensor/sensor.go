package sensor

type Sensor struct {
	Threshhold float32
	Thermo     Thermo
	SL         SignalL
}

func NewSensor(threshhold float32, thermo Thermo, SL SignalL) *Sensor {
	return &Sensor{
		Threshhold: threshhold,
		Thermo:     thermo,
		SL:         SL,
	}
}

func (s *Sensor) GetCurrentTemperature() float32 {
	return s.Thermo.GetTemperature()
}

func (s *Sensor) CheckThreshhold(temp float32) {
	if temp <= s.Threshhold {
		s.SL.TurnOn()
	}
}

type Thermo interface {
	GetTemperature() float32
}

//go:generate mockery --output=./mocks --name=SignalL
type SignalL interface {
	TurnOn()
}
