package thermo

type Thermo struct {
	Temp float32
}

func NewThermo(temp float32) *Thermo {
	return &Thermo{
		Temp: temp,
	}
}

func (t *Thermo) SetTemperature(a float32) {
	t.Temp = a
}

func (t *Thermo) GetTemperature() float32 {
	return t.Temp
}
