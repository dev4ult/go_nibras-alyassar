package main

import (
	"fmt"
)

type Car struct {
	CarType string
	FuelIn  float64
}

func (c Car) getMaxDistance() string {
	distance := c.FuelIn * 1.5
	return "car type: " + c.CarType + ", est. distance: " + fmt.Sprintf("%.2f", distance)
}

func main() {
	suvCar := Car{
		CarType: "SUV",
		FuelIn:  10.5,
	}

	fmt.Println(suvCar.getMaxDistance())
}