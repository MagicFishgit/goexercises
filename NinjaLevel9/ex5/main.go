package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var incr uint64 //incrementor

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			atomic.AddUint64(&incr, 1)
			fmt.Println("Increment is:", atomic.LoadUint64(&incr))
			wg.Done()

		}()
	}

	wg.Wait()
	println("All routines are done")
	println("Value of incrementor is:", incr, "and should be 100 if no race conditions")
	println("Value is 100 therefor there is no race condition. Run 'go run -race main.go' to test for race condition.")

}
