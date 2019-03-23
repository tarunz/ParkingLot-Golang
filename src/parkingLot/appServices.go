package parkinglot

import "fmt"

// NewParkingLot initializes parking lot with given capacity and returns reference to it.
func NewParkingLot(capacity int) *ParkingLot {
	fmt.Printf("Created a parking lot with %d slots\n", capacity)
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

func (p *ParkingLot) Leave(slotNo int) {
	fmt.Printf("Slot number %d is free\n", slotNo)
	p.Slots[slotNo-1].freeSlot()
	p.FreeSlots++
}

func (p ParkingLot) Status() {
	fmt.Printf("Slot No.\tRegistration No\t\tColour\n")
	n := p.Capacity
	regNo := ""
	var color Color
	for i := 0; i < n; i++ {
		if p.Slots[i].Occupied {
			regNo, color = p.Slots[i].vehicleDetails()
			fmt.Printf("%d\t\t%s\t\t%s\n", i+1, regNo, color)
		}
	}
}

func (p ParkingLot) GetRegNoOfColor(color string) {
	c := StringToColor(color)
	n := p.Capacity
	firstEntry := true
	for i := 0; i < n; i++ {
		if p.Slots[i].Occupied {
			carr, carc := p.Slots[i].vehicleDetails()
			if c == carc {
				if firstEntry {
					fmt.Printf("%s", carr)
					firstEntry = false
				} else {
					fmt.Printf(", %s", carr)
				}
			}
		}
	}
	if !firstEntry {
		fmt.Printf("\n")
	}
}

func (p ParkingLot) GetSlotNoOfColor(color string) {
	c := StringToColor(color)
	n := p.Capacity
	firstEntry := true
	for i := 0; i < n; i++ {
		if p.Slots[i].Occupied {
			_, carc := p.Slots[i].vehicleDetails()
			if c == carc {
				if firstEntry {
					fmt.Printf("%d", i+1)
					firstEntry = false
				} else {
					fmt.Printf(", %d", i+1)
				}
			}
		}
	}
	if !firstEntry {
		fmt.Printf("\n")
	}
}

func (p ParkingLot) GetSlotNoOfCar(regNo string) {
	n := p.Capacity
	for i := 0; i < n; i++ {
		if p.Slots[i].Occupied {
			carr, _ := p.Slots[i].vehicleDetails()
			if regNo == carr {
				fmt.Printf("%d\n", i+1)
				return
			}
		}
	}
	fmt.Printf("Not found\n")
}
