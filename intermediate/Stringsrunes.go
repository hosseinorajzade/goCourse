package main

import "fmt"

func main() {
	message := "Hello, \nworld"
	message2 := "Hello, \tworld"
	message3 := "Hello \rworld"
	rawMessage := `Hello\nworld`
	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)
	greetin := "Hello"
	name := "hossein"
	fmt.Println(greetin + " " + name)

}
