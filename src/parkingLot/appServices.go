package parkinglot

import "fmt"

// NewParkingLot initializes parking lot with given capacity and returns reference to it.
func NewParkingLot(capacity int) *ParkingLot {
	if capacity <= 0 {
		fmt.Printf("No of slots in lot should be non negative.")
		return nil
	}
	fmt.Printf("Created a parking lot with %d slots\n", capacity)
	return &ParkingLot{capacity, make([]Slot, capacity), capacity}
}

// MakeParkingLot initilizes parking lot with given capacity and returns it.
func MakeParkingLot(capacity int) ParkingLot {
	return ParkingLot{capacity, make([]Slot, capacity), capacity}
}

// Park takes regNo and color of Car and parks car if slot is available.
func (p *ParkingLot) Park(regNo string, color string) {
	if p == nil {
		fmt.Println("Parking lot does not exixts.")
		return
	}
	slotNo := p.availableSlot()
	if slotNo == -1 {
		fmt.Printf("Sorry, parking lot is full\n")
		return
	}
	fmt.Printf("Allocated slot number: %d\n", slotNo+1)
	p.parkCar(MakeCar(regNo, StringToColor(color)), slotNo)
}

// Leave takes slot of which car is leaving and marks it free.
func (p *ParkingLot) Leave(slotNo int) {
	if p == nil {
		fmt.Println("Parking lot does not exixts.")
		return
	}
	if slotNo < 1 || slotNo > p.Capacity {
		fmt.Println("Invalid slot number.")
	}
	fmt.Printf("Slot number %d is free\n", slotNo)
	p.Slots[slotNo-1].freeSlot()
	p.FreeSlots++
}

// Status prints the summary of all the vehicles in Parking lot.
func (p ParkingLot) Status() {
	fmt.Printf("Slot No.    Registration No    Colour\n")
	n := p.Capacity
	regNo := ""
	var color Color
	for i := 0; i < n; i++ {
		if p.Slots[i].Occupied {
			regNo, color = p.Slots[i].vehicleDetails()
			fmt.Printf("%d           %s      %s\n", i+1, regNo, color)
		}
	}
}

// GetRegNoOfColor take color as input and prints all the cars of that color in parking lot.
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

//	GetSlotNoOfColor take color as input and prints all the slots, cars of which is of that color in parking lot.
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

// GetSlotNoOfCar prints the slotNo of car with regNo(input)
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
