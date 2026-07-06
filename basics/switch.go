package main

import "fmt"

func main() {
	fruit := "pineapple"
	switch fruit {
	case "apple":
		fmt.Println("it,s apple")
	case "banana":
		fmt.Println("it,s banana")

	default:
		fmt.Println("Unkown fruit")
	}

	day := "3shanbe"
	switch day {
	case "shanbe", "1shanbe", "2shanbe", "3shanbe", "4shanbe":
		fmt.Println("its week")
	case "5shanbe", "jome":
		fmt.Println("it weekend")
	default:
		fmt.Println("Unkwon day")
	}
}
