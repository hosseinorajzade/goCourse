package main

import "fmt"

func main() {
	originalArray := [3]int{0, 1, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100
	fmt.Println(originalArray)
	// for i, v := range originalArray {
	// 	fmt.Printf("index:%d , value:%v\n", i, v)
	// }

	// a, _ := somereturn()
	// fmt.Println(a)
	// // fmt.Println()
	// fmt.Println("The length is ", len(originalArray))
	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{1, 2, 3}
	// fmt.Println("Array1 is equal Array2: ", array1 == array2)
}

func somereturn() (int, int) {
	return 1, 2
}
