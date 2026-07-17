package main

import "fmt"

func main() {
	s := fmt.Sprint("Hello", "world", 123, 456)
	fmt.Print(s)
	var name string
	var age int
	fmt.Print("enter your name:")
	fmt.Scan(&name)
	fmt.Print("enter your age:")
	fmt.Scan(&age)

	fmt.Print(fmt.Sprint(name, age))

}
