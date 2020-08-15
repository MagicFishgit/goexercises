package main

//imports
import "fmt"

//package scope vars
type magicfish int

var x magicfish

//main func
func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
