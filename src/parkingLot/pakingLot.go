package parkinglot

// ParkingLot struct which has list of Slot and total number of free Slots
type ParkingLot struct {
	Capacity  int
	Slots     []Slot
	FreeSlots int
}

// NewParkingLot initializes parking lot with given capacity and returns reference to it.
func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{capacity, make([]Slot, capacity), capacity}
}

// MakeParkingLot initilizes parking lot with given capacity and returns it.
func MakeParkingLot(capacity int) ParkingLot {
	return ParkingLot{capacity, make([]Slot, capacity), capacity}
}
