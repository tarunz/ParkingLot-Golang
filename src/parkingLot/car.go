package parkinglot

// A Car struct containing registration no. and its color
type Car struct {
	iVehicle
	regNo string
	color Color
}

// NewCar creates new Car and returns its reference.
func NewCar(regNo string, color Color) *Car {
	c := Car{}
	c.regNo = regNo
	c.color = color
	return &c
}

// MakeCar creates new Car and returns it.
func MakeCar(regNo string, color Color) Car {
	c := Car{}
	c.regNo = regNo
	c.color = color
	return c
}

// SetRegNo for setting Reg No. of Car
func (v *Car) SetRegNo(regNo string) {
	v.regNo = regNo
}

// GetRegNo for getting Reg No. of Car
func (v Car) GetRegNo() string {
	return v.regNo
}

// SetColor for setting Color of Car
func (v *Car) SetColor(color Color) {
	v.color = color
}

// GetColor for getting Color of Car
func (v Car) GetColor() Color {
	return v.color
}
