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
			sum = sum * count
			count--
		}
	}
	return sum
}
