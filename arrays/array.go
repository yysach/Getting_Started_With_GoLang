package arrays

import "fmt"

func Array() {
	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Printing Array: ", array)

	var prefixSum int
	for i := 0; i < len(array); i++ {
		prefixSum += array[i]
	}
	fmt.Println("Printing prefixSum: ", prefixSum)

	// multiDimensional array
	var _2D = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var i, j int
	for i = 0; i < len(_2D); i++ {
		for j = 0; j < len(_2D[0]); j++ {
			fmt.Print(_2D[i][j], " ")
		}
		fmt.Print("\n")
	}

	//slicing in go like python
	var sliced []int = array[2:4]
	fmt.Println("Printing sliced: ", sliced)
}
