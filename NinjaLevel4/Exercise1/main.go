package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [5]int

	for i := 0; i < 5; i++ {
		a[i] = i + rand.Intn(100)
	}

	for i, v := range a {
		fmt.Printf("%T\t%v\t\n", i, v)
	}

}
