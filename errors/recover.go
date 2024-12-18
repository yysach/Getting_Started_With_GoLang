package errors

/*
Recover is used to regain control of a program from panic or error-condition.
It stops the terminating sequence and resumes the normal execution.
It is called from the deferred function. It retrieves the error value passed through the call of panic.
Normally, it returns nil which has no other effect.
*/
import (
	"fmt"
)

func recovery() {
	fmt.Println("Result: ", SaveDivide(10, 0))
	fmt.Println("Result: ", SaveDivide(10, 10))
}
func SaveDivide(num1, num2 int) int {
	defer func() {
		fmt.Println("Recovery(): ", recover())
	}()
	quotient := num1 / num2
	return quotient
}
