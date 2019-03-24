package parkinglot

import "testing"

func TestNewParkingLotInputCapacity(t *testing.T) {
	input := 0
	actualResult := NewParkingLot(input)
	if actualResult != nil {
		t.Fatalf("Expected %s but got %s", nil, actualResult)
	}
}

func TestPark(t *testing.T) {
	p := NewParkingLot(1)
	actualResult := p.availableSlot()
	expectedResult := 0
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
	p.Park("123", "Black")
	actualResult = p.availableSlot()
	expectedResult = -1
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestLeave(t *testing.T) {
	p := NewParkingLot(1)
	p.Park("123", "Black")
	p.Leave(1)
	actualResult := p.availableSlot()
	expectedResult := 0
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestGetSlotNoOfCar(t *testing.T) {
	p := NewParkingLot(1)
	p.Park("123", "Black")
}

func TestColorToStringColorInput(t *testing.T) {
	c := Black
	var expectedResult = "Black"
	var actualResult = c.String()
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestStringToColorStringInput(t *testing.T) {
	c := "Black"
	var expectedResult = Black
	var actualResult = StringToColor(c)
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
