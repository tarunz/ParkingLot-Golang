package parkinglot

import "fmt"

// ParkingLot struct which has list of Slot and total number of free Slots
type ParkingLot struct {
	Capacity  int
	Slots     []Slot
	FreeSlots int
}

// NewParkingLot initializes parking lot with given capacity and returns reference to it.
func NewParkingLot(capacity int) *ParkingLot {
	fmt.Printf("Created a parking lot with %d slots", capacity)
	return &ParkingLot{capacity, make([]Slot, capacity), capacity}
}

// MakeParkingLot initilizes parking lot with given capacity and returns it.
func MakeParkingLot(capacity int) ParkingLot {
	return ParkingLot{capacity, make([]Slot, capacity), capacity}
}

func (p *ParkingLot) Park(regNo string, color string) {
	slotNo := p.availableSlot()
	if slotNo == -1 {
		fmt.Printf("Sorry, parking lot is full\n")
		return
	}
	fmt.Printf("Allocated slot number: %d\n", slotNo+1)
	p.parkCar(MakeCar(regNo, StringToColor(color)), slotNo)
}

func (p *ParkingLot) parkCar(c Car, i int) {
	p.Slots[i].carToslot(c)
	p.FreeSlots--
}

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

func (p *ParkingLot) Leave(slotNo int) {
	fmt.Printf("Slot number %d is free\n", slotNo)
	p.Slots[slotNo-1].freeSlot()
	p.FreeSlots++
}

func (p ParkingLot) Status() {
	fmt.Printf("Slot No.\tRegistration No\t\tColour\n")
	n := p.Capacity
	regNo := ""
	color := Black
	for i := 0; i < n; i++ {
		if p.Slots[i].Occupied {
			regNo, color = p.Slots[i].vehicleDetails()
			fmt.Printf("%d\t\t%s\t\t%s\n", i+1, regNo, color.ColorToString())
		}
	}
}

func (p ParkingLot) GetRegNoOfColor(color string) {
	c := StringToColor(color)
	n := p.Capacity
	for i := 0; i < n; i++ {
		if p.Slots[i].Occupied {
			carr, carc := p.Slots[i].vehicleDetails()
			if c == carc {
				fmt.Printf("%s,", carr)
			}
		}
	}
}
