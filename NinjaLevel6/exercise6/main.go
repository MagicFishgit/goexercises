package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is an anonamous function using immediately invokable execution I learnt from js")
	}()
}
