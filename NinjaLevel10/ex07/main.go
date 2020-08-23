package main

import "fmt"

var c chan int

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			genRoutine()
		}
	}()

	for v := range c {
		fmt.Println(v)
	}

}

func genRoutine() {

	go func(c chan<- int) {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}(c)

}

//hmmm didn't work this way... will return to this.
