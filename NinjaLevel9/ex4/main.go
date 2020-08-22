package main

import (
	"fmt"
	"sync"
)

var incr int //incrementor
var mux sync.Mutex

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			mux.Lock() //locking the vars used here.
			counter := incr
			counter++
			incr = counter
			fmt.Println("Increment is:", incr)
			mux.Unlock() //unlocking the vars for the next goroutine to use since I am done with it.
			wg.Done()

		}()
	}

	wg.Wait()
	println("All routines are done")
	println("Value of incrementor is:", incr, "and should be 100 if no race conditions")
	println("Value is 100 therefor there is no race condition. Run 'go run -race main.go' to test for race condition.")

}
