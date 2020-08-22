package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incr int //incrementor

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			counter := incr
			runtime.Gosched()
			counter++
			incr = counter
			fmt.Println("Increment is:", incr)
			wg.Done()
		}()
	}

	wg.Wait()
	println("All routines are done")
	println("Value of incrementor is:", incr, "and should be 100 if no race conditions")
	println("Value is not 100 therefor there is a race condition. Run 'go run -race main.go' to test for race condition.")

}
