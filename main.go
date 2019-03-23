package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	pk "./src/parkingLot"
)

func main() {

	arg := os.Args[1:]

	if len(arg) > 0 {
		// File Mode
		file, err := os.Open(arg[0])
		if err != nil {
			fmt.Printf("File does not exist: %s\n", arg[0])
		} else {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				// fmt.Println(scanner.Text())
				processString(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
		}
		defer file.Close()
	} else {
		// Interactive Mode
		for {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := scanner.Text()
			if text == "exit" {
				break
			}
			processString(text)
		}
	}
}

var plot *pk.ParkingLot

func processString(s string) {
	vals := strings.Split(s, " ")
	switch vals[0] {
	case "create_parking_lot":
		n, err := strconv.Atoi(vals[1])
		if err != nil {
			fmt.Println("CREATE: Invalid value for number of slots.")
		} else {
			plot = pk.NewParkingLot(n)
		}
	case "park":
		plot.Park(vals[1], vals[2])
	case "leave":
		i, err := strconv.Atoi(vals[1])
		if err != nil {
			fmt.Println("LEAVE: Invalid value for slot number.")
		} else {
			plot.Leave(i)
		}
	case "status":
		plot.Status()
	case "registration_numbers_for_cars_with_colour":
		plot.GetRegNoOfColor(vals[1])
	case "slot_numbers_for_cars_with_colour":
		plot.GetSlotNoOfColor(vals[1])
	case "slot_number_for_registration_number":
		plot.GetSlotNoOfCar(vals[1])
	default:
		fmt.Println("ACTION: Invalid action")
	}
}
