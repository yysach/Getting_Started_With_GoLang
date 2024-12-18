package construct_types

import "fmt"

/*
	A pointer is a variable that stores the address of another variable.
	var var_name *var_type
	A newly declared pointer which has not been assigned to a variable has the nil value.
	The address-of operation &, when placed before a variable gives us the memory address of variable
	With pointers, we can pass a reference to a variable, instead of passing a copy of the variable which can reduce memory usage
*/

func swapByValue(a, b int) {
	a, b = b, a
	fmt.Println("Inside swapByValue - a:", a, "b:", b) // Prints swapped values
}

func swapByReference(a, b *int) {
	// Swapping values using dereferencing pointers
	*a, *b = *b, *a
	fmt.Println("Inside swapByReference - a:", *a, "b:", *b) // Prints swapped values
}
