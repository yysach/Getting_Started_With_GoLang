package functions

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
}

// func(arg1 type) functionName returnType { body }
func (e Employee) fullName() string { return e.firstName + " " + e.lastName }

func Function() {
	e1 := Employee{firstName: "John", lastName: "Doe"}
	fmt.Println("Full Name of employee :", e1.fullName())

	// anonymous function
	number := 10
	squareNum := func() int {
		return number * number
	}
	fmt.Println("Square of ", number, "is :", squareNum())

	// computing factorial of 5
	var num int = 5
	fmt.Println("Factorial of ", num, "is :", factorial(num))
}
