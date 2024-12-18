package data_types

import "fmt"

func DataType() {
	var currentYear int = 2023
	var amount float32 = 9.9
	var flag bool = true
	var name string = "sachin"
	var balance float64 = 2234.8034

	// printing variable using println
	fmt.Println("current year: ", currentYear)
	fmt.Println("amount Paid: ", amount)
	fmt.Println("flag: ", flag)
	fmt.Println("name: ", name)
	fmt.Println("balance: ", balance)

	// printing variable using printf
	// printing data types
	fmt.Printf("%T %T %T %T %T\n", currentYear, amount, flag, name, balance)
	fmt.Println("printing all variable values: ", currentYear, amount, flag, name, balance)
	fmt.Println("calling trial function in same module...")
	trial()
}
