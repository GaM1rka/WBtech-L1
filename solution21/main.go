package main

import (
	"fmt"
)

type Liters float64
type Gallons float64

// Sturcture that will be adapted
type Car struct {
	fuel Liters
}

func (c Car) GetFuel() Liters {
	return c.fuel
}

// The main interface
type FuelChecker interface {
	Check() bool
}

// Adapter that implement our interface
type FuelAdapter struct {
	car *Car
}

func NewFuelAdapter(car *Car) *FuelAdapter {
	return &FuelAdapter{car: car}
}

func (a *FuelAdapter) Check() bool {
	// Converting liters to gallons
	gallons := Gallons(a.car.GetFuel() * 0.264)
	return gallons > 5
}

func enough(fc FuelChecker) bool {
	return fc.Check()
}

func main() {
	car := &Car{fuel: 20}

	adapter := NewFuelAdapter(car)

	fmt.Println("Fuel enough: ", enough(adapter))
}
