package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	x := person{
		first: "Rudi",
		last:  "Visagie",
		age:   25,
	}

	x.speak()

}

func (p person) speak() {

	fmt.Println("My Name is:", p.first, p.last, "and I am ", p.age, " years old")
}
