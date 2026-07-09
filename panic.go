package main

import "fmt"

func main() {
	process(-1)
}

func process(input int) {
	if input < 0 {
		panic("khaye kardam haji")
	}
	fmt.Println("Its ok")
}
