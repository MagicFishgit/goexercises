package main

import "fmt"

func main() {
	x := 100
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\t", y, y, y)
}
