package main

import "fmt"

func main() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"one", "two", "three"}

	fmt.Println(x1)
	fmt.Println(x2)

	xx12 := [][]string{x1, x2}

	fmt.Println(xx12)

	for i, xs := range xx12 {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
