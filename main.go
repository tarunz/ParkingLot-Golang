package main

import (
	"fmt"

	pk "./src/parkinglot"
)

func main() {
	// c := pk.NewCar("123", "Black")
	// c.SetColor("G")

	p := pk.NewParkingLot(6)
	fmt.Println(p.Cars)
}
