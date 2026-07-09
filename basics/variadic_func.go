package main

import "fmt"

func main() {

	fmt.Println("sum 1 , 2 ,3: ", sum(1, 2, 3))
}
func sum(numbs ...int) int {
	total := 0
	for _, v := range numbs {
		total += v
	}
	return total
}
