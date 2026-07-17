package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}
func main() {

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println(area)
	rect.Scale(2)
	area = rect.Area()

	fmt.Println(area)
	num := MyInt(-5)
	num1 := MyInt(5)

	fmt.Println(num.isPositive(), num1.isPositive())
	fmt.Println(num.WelcomeMessage())
}

type MyInt int

func (m MyInt) isPositive() bool {
	return m > 0
}
func (MyInt) WelcomeMessage() string {
	return "Welcome to My App"
}
