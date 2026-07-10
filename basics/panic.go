package main

import "fmt"

func main() {
	process(-1)
}

func process(input int) {
	if input < 0 {
		panic("its not OK")
	}
	fmt.Println("Its ok")
}
