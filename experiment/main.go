package main

import "fmt"

//functions

func testfunc(t string, t2 int) {
	println(t, t2)
}

func returnfunc(s string) string {
	return fmt.Sprint("Hello ", s)
}

func multiReturnFunc(s string) (string, bool) {
	return fmt.Sprint("Hello ", s), true
}

func variadic(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T", x)

	sum := 0

	for i, v := range x {
		sum += v
		println("index ", i, " adding to the sum: ", v, " Total: ", sum)
	}
}

func main() {
	testfunc("test", 2)
	s := returnfunc("Rudi")
	println(s)
	s2, s3 := multiReturnFunc("Sas")
	println(s2, s3)
	variadic(1, 2, 3, 4, 5, 6, 7, 7, 8, 6, 5, 3, 33, 4, 254, 256, 36, 34, 52, 42, 4, 24, 23, 423, 4, 235, 34, 63, 6, 3) //variadic.. can pass unlimited number of arguments of type int.
}

//testing functions and displaying multiple ways of implimenting them
