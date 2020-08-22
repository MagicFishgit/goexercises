package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//if this func had a non pointer receiver  ie. "person"  then I could pass in a pointer value &person or just normal type value person
func (m *person) speak() { //points to the address of person - pointer receiver on method speak can only accept pointer value p1
	fmt.Println("Hi I am", m.first, m.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := &person{ //value of address space - pointer value which can only go to pointer receiver.
		first: "Rudi",
		last:  "Visagie",
		age:   25,
	}

	saySomething(p1)

}
