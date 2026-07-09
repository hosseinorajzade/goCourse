package main

import "fmt"

func main() {

	fmt.Println(add(1, 2))
	operation := add
	result := operation(3, 5)
	fmt.Println(result)
	fmt.Println(applyoperation(5, 10, add))
	multiplierby2 := multiplier(2)

	fmt.Println(multiplierby2(6))

}

func add(a int, b int) int {

	return a + b
}

func applyoperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}
func multiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}
