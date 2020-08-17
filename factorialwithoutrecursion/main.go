package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(num int) int {

	sum := num
	count := num - 1

	for i := 0; i < num; i++ {

		if count != 0 {
			sum *= count
			count--
		}
	}
	return sum
}

//challenge from Todd to do a factorial without using recursion and using loops.
