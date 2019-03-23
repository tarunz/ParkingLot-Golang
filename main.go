package main

import (
	pk "./src/parkinglot"
)

func main() {
	// c := pk.NewCar("123", "Black")
	// c.SetColor("G")

	p := pk.NewParkingLot(6)
	// p.Cars[2] = pk.MakeCar("123", pk.Orange)
	// p.Park("123", "Black", 1)
	p.Park("KA-01-HH-1234", "White")
	p.Park("KA-01-HH-9999", "White")
	p.Park("KA-01-BB-0001", "Black")
	p.Park("KA-01-HH-7777", "Red")
	p.Park("KA-01-HH-2701", "Blue")
	p.Park("KA-01-HH-3141", "Black")
	p.Leave(4)
	p.Status()
	p.Park("KA-01-P-333", "White")
	p.Park("DL-12-AA-9999", "White")
	p.GetRegNoOfColor("White")
	p.GetSlotNoOfColor("White")
	p.GetSlotNoOfCar("KA-01-HH-3141")
	p.GetSlotNoOfCar("MH-04-AY-1111")
	// fmt.Println(p)
	// p.Leave(1)
	// fmt.Println(p)
}
