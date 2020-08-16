package main

import "fmt"

//type declaration
type magicfish struct {
	animal  bool
	magical bool
}

func (z magicfish) greet() {

	fmt.Println("Hello ", z.animal, " you have tested ", z.magical, " for magical ability.") //whatever
}

func main() {
	sub1 := magicfish{
		animal:  true,
		magical: true,
	}

	sub1.greet()
}
