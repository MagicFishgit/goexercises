package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	d := make(chan int, 1)

	d <- 43

	fmt.Println("buffered value: ", <-d)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
