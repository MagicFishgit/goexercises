package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hey im running something")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hey, me too")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Hey the groups are done waiting yay.")

}
