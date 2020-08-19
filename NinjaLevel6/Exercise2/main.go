package main

import "fmt"

func main() {

	xints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 7, 8, 6, 5, 4, 4453, 5, 34, 534, 534, 534, 534, 5}

	fmt.Println("The summed amount is: ", foo(xints...))
	fmt.Println("Bar output: ", bar(xints))

}

func foo(a ...int) int {

	var sum int

	for _, j := range a {
		sum += j
	}

	return sum

}

func bar(a []int) int {

	var sum int

	for _, j := range a {
		sum += j
	}

	return sum

}
