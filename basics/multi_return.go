package main

import (
	"errors"
	"fmt"
)

func main() {
	q, r := divide(1, 1)
	fmt.Printf("first value %d , second value %d\n", q, r)

	result, err := compare(2, 2)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Result", result)
	}
}

func divide(a int, b int) (c int, d int) {

	c = a / b
	d = a % b
	return

}

func compare(a, b int) (string, error) {

	if a > b {
		return " a is bigger than b", nil
	} else if a < b {
		return "b is bigger than a", nil
	} else {
		return "", errors.New("Unable to compare")
	}

}
