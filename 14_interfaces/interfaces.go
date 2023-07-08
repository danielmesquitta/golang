//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package interfaces

import "fmt"

type Lift byte

const (
	SmallLift Lift = iota
	StandardLift
	LargeLift
)

func (l Lift) String() string {
	switch l {
	case SmallLift:
		return "small lift"
	case StandardLift:
		return "standard lift"
	case LargeLift:
		return "large lift"
	default:
		return "unknown"
	}
}

type Driver interface {
	Drive(to Lift)
}

type Car string

func (c Car) Drive(to Lift) {
	fmt.Printf("Driving a %s car to a %s\n", c, to)
}

type Truck string

func (t Truck) Drive(to Lift) {
	fmt.Printf("Driving a %s truck to a %s\n", t, to)
}

type Motorcycle string

func (m Motorcycle) Drive(to Lift) {
	fmt.Printf("Driving a %s motorcycle to a %s\n", m, to)
}

func driveVehicle(d Driver, to Lift) {
	d.Drive(to)
}

func Interfaces() {
	car, truck, motorcycle := Car("Road Devourer"), Truck("Big Truck"), Motorcycle("Small Motorcycle")

	driveVehicle(car, StandardLift)

	driveVehicle(truck, LargeLift)

	driveVehicle(motorcycle, SmallLift)
}
