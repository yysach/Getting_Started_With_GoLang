package controls

import (
	"fmt"
	"strconv"
)

func Control() {
	var age int
	fmt.Println("Enter age :")
	fmt.Scanln(&age)

	// if else
	if age >= 18 {
		fmt.Println("Eligible for casting vote !!")
	} else {
		fmt.Println("Not eligible for casting vote !!")
	}

	// nested for loop
	for a := 0; a < 2; a++ {
		for b := 2; b > 0; b-- {
			fmt.Print(a, " ", b, "\n")
		}
	}

	// for each iteration
	nums := []int{2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Total Sum :", sum)

	// defining go constant
	const myName string = "SachinY"
	fmt.Printf("My Name is %s \n", myName)

	// type casting
	var i int = 10
	var f float64 = 6.44
	var str1 string = "101"
	var str2 string = "12.123"

	// numerical to numerical type casting
	fmt.Println(float64(i))
	fmt.Println(int(f))

	// string to numerical typecasting
	newInt, _ := strconv.ParseInt(str1, 0, 64)
	fmt.Println(newInt)

	balance, _ := strconv.ParseFloat(str2, 64)
	fmt.Println(balance)
}
