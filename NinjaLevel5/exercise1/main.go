package main

import "fmt"

//type declaration
type person struct {
	first        string
	last         string
	iceCreamFlav []string
}

func main() {

	p1 := person{
		first:        "Rudi",
		last:         "Visagie",
		iceCreamFlav: []string{"Bubblegum", "Vanilla", "Bannana"},
	}

	p2 := person{
		first:        "Sas",
		last:         "Swart",
		iceCreamFlav: []string{"Anxiety", "Ignoring", "Oblivious to Emotions"},
	}

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.iceCreamFlav {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.iceCreamFlav {
		fmt.Println(i, v)
	}
}
