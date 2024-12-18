package construct_types

import (
	"fmt"
	"reflect"
)

func ConstructType() {
	fmt.Println("Diff construct in Go")
	structType()
	interfaceRunner()

	// call by value and call by reference in Go
	// pointer
	var x, y = 10, 20
	fmt.Println("Before swapByValue - x:", x, "y:", y)
	swapByValue(x, y)                                 // Call by value
	fmt.Println("After swapByValue - x:", x, "y:", y) // Original values are unchanged

	fmt.Println("Before swapByReference - x:", x, "y:", y)
	swapByReference(&x, &y)                               // Call by reference (passing memory addresses)
	fmt.Println("After swapByReference - x:", x, "y:", y) // Values are swapped

	//Go reflect
	/*
		Go reflection is the ability of a program to examine its own structure, particularly through type.
		It is a form of meta programming
	*/

	age := 27.5
	fmt.Printf("Meta programming of age %T\n", age)
	fmt.Println(reflect.TypeOf(age))
}
