package main

import "fmt"

func main() {

	fmt.Println("whooo im running first")
	defer deferfunc()
	fmt.Println("I just called defer on a function which will run after me cause I am the last function in main function.")
}

func deferfunc() {
	fmt.Println("This runs after main is done cause it is defferred")
}
