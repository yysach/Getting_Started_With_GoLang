package errors

import (
	"errors"
	"fmt"
	"math"
)

/*
	Go does not have exception mechanism like try/catch in Java, we cannot throw an exception in Go.
	Go uses different mechanism which is know as defer-panic-and-recovery mechanism.
	We should always check the error at the calling statement, if we receive any of it or not.
	We should never ignore errors, it may lead to program crashes.

The way go detect and report the error condition is
	A function which can result in an error returns two variables: a value and an error-code which is nil in case of success, and != nil in case of an error-condition.
	The error is checked, after the function call . In case of an error ( if error != nil), the execution of the actual function (or if necessary the entire program) is stopped.

	Go has predefined error interface type
	type error interface {
		Error() string
	}

	err := errors.New("math - square root of negative number")
*/

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

func Errors() {
	result, err := Sqrt(-64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result, err = Sqrt(64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	recovery()
	defers()
	panicking()
}
