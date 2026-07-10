package main

import "fmt"

func main() {
	process()
	fmt.Println("Returned from procces")
}
func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered", r)
		}

	}()

	fmt.Println("strat procces")
	panic("Something went wrong")
	fmt.Println("EndProcces")
}
