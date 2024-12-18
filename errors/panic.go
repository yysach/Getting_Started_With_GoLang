package errors

/*
Go panic is a mechanism by which we handle error situations. Panic can be used to abort a function execution. When a function calls panic, its execution stops and the control flows to the associated deferred function.

The caller of this function also gets terminated and caller's deferred function gets executed (if present any). This process continues till the program terminates. Now the error condition is reported.

This termination sequences is called panicking and can be controlled by the built-in function recover.
*/
import "fmt"

func panicking() {
	fmt.Println("Calling x from main.")
	x()
	fmt.Println("Returned from x.")
}
func x() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in x", r)
		}
	}()
	fmt.Println("Executing x...")
	fmt.Println("Calling y.")
	y(0)
	fmt.Println("Returned normally from y.")
}

func y(i int) {
	fmt.Println("Executing y....")
	if i > 2 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in y", i)
	fmt.Println("Printing in y", i)
	y(i + 1)
}
