package main

import (
	"fmt"
)

func main() {
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	ftm.Println("Index %d , value %d", index, value)
	// }
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ODD is = %d ", i)
		if i == 5 {
			break
		}
	}
}
