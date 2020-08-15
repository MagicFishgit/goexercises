package main

import "fmt"

func main() {
	const (
		a     = 69 //UNTYPED
		b int = 70 //TYPED
	)

	fmt.Println(a, b)
}
