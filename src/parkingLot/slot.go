package parkinglot

type Slot struct {
	Occupied bool
	Car      Car
}

func (s *Slot) carToslot(c Car) {
	s.Occupied = true
	s.Car = c
}

func (s *Slot) freeSlot() {
	s.Occupied = false
}

func (s Slot) vehicleDetails() (regNo string, color Color) {
	return s.Car.GetRegNo(), s.Car.GetColor()
}
