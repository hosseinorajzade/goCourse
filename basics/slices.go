package basics

import (
	"fmt"
)

func main() {
	// slice := make([]int, 5)
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]
	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("Slice1 :", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println(sliceCopy)
	fmt.Println("Slice copy", sliceCopy)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("adding Value %d in outter slice at index %d and inner slice at inde %d\n", i+j, i, j)
		}

	}
	fmt.Println("twoD", twoD)

	slice2 := slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("slice2 cap is ", cap(slice2))
	fmt.Println("slice2 len is ", len(slice2))

}
