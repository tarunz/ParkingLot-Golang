package parkinglot

// ParkingLot struct which has list of Slot and total number of free Slots
type ParkingLot struct {
	Capacity  int
	Slots     []Slot
	FreeSlots int
}

// Takes car and slotNo. as input and parks car in it.
func (p *ParkingLot) parkCar(c Car, i int) {
	p.Slots[i].carToslot(c)
	p.FreeSlots--
}

// Returns lowest indexed slotNo. if any, otherwise returns -1
func (p ParkingLot) availableSlot() int {
	if p.FreeSlots > 0 {
		n := p.Capacity
		for i := 0; i < n; i++ {
			if !p.Slots[i].Occupied {
				return i
			}
		}
	}
	return -1
}
