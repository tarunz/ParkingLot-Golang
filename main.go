package main

import (
	"fmt"

	pk "./src/parkinglot"
)

func main() {
	// c := pk.NewCar("123", "Black")
	// c.SetColor("G")

	p := pk.NewParkingLot(6)
	p.Cars[2] = pk.MakeCar("123", pk.Orange)
	fmt.Println(p.Cars)
}
