package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
	address   Address
}

type Address struct {
	city    string
	country string
}

func main() {

	p := Person{
		firstname: "Hossein",
		lastname:  "Orajzade",
		age:       27,
		address: Address{
			city:    "esfahan",
			country: "iran",
		},
	}

	p1 := Person{
		firstname: "Matin",
		lastname:  "fathi",
		age:       26,
		address: Address{
			city:    "London",
			country: "U.k",
		},
	}

	fmt.Println(p.fullname())
	fmt.Println(p.address.city)
	fmt.Println(p1.firstname)
	fmt.Println(p1.address.city)
	fmt.Println("Before increament age", p.age)
	p.increamentAge()
	fmt.Println("After increament age", p.age)

}

func (p Person) fullname() string {
	return p.firstname + " " + p.lastname
}

func (p *Person) increamentAge() {
	p.age++
}
