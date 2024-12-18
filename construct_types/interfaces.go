package construct_types

import (
	"fmt"
)

/*
	Generally, the name of an interface is formed by the method name plus the [e]r suffix, such as Printer, Reader, Writer, Logger, Converter, etc.

A type doesn't have to state explicitly that it implements an interface: interfaces are satisfied implicitly. Multiple types can implement the same interface.
A type that implements an interface can also have other functions.
A type can implement many interfaces.
An interface type can contain a reference to an instance of any of the types that implement the interface
*/
type Vehicle interface {
	accelerate()
}

func printVehicle(v Vehicle) {
	fmt.Println("Vehicle: ", v)
}

type car struct {
	model string
	color string
}

func (c car) accelerate() {
	fmt.Println("Accelerate: car", c.model)
}

type bike struct {
	model string
	color string
	speed int
}

func (b bike) accelerate() {
	fmt.Println("Accelerate: bike", b.model)
}

func VehicleFactory(modelType string) Vehicle {
	switch modelType {
	case "car":
		return car{"suzuki", "blue"}
	case "bike":
		return bike{"pulsar", "Red", 200}
	default:
		fmt.Println("Vehicle Factory Not Found")
		return nil
	}
}

func interfaceRunner() {
	c1 := VehicleFactory("car")
	b1 := VehicleFactory("bike")
	c1.accelerate()
	b1.accelerate()
	printVehicle(c1)
	printVehicle(b1)
}
