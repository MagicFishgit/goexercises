package main

import "fmt"

func main() {

	funkyfunc(func() {
		fmt.Println("gosh I need to get past this basic stuff...")
	})
}

func funkyfunc(a func()) {
	a()
}
