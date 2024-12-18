package main

import (
	"fmt"
	"hello/arrays"
	"hello/construct_types"
	"hello/controls"
	"hello/data_types"
	"hello/errors"
	"hello/functions"
	"hello/maps"
	"os"
)

func main() {
	fmt.Println("Hello, World! ######### Main")
	// if we declare identifier in lowercase letter, it will be visible within package only.
	// But if we declare package in uppercase letter, it will be visible within and outside the package which is also knows as exported
	data_types.DataType()

	controls.Control()

	functions.Function()

	arrays.Array()

	// command line argument
	/*
		The os.Args is used to get the arguments. The index 0 of os.Args contains the path of the program
		The os.Args[1:] hold provided arguments
	*/
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}
	fmt.Println(s)

	// construct like struct
	/*
		Go has different approaches to implement the concepts of object-orientation.
		Go does not have classes and inheritance. Go fulfill these requirements through its powerful interface
		An interface defines a set of abstract methods and does not contain any variable

	*/
	construct_types.ConstructType()

	maps.Maps()

	errors.Errors()
}
