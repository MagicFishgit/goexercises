package main

import "fmt"

//functions

func testfunc(t string, t2 int) {
	println(t, t2)
}

func returnfunc(s string) string {
	return fmt.Sprint("Hello ", s)
}

func main() {
	testfunc("test", 2)
	s := returnfunc("Rudi")
	println(s)
}

//testing functions
