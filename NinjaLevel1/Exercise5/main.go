package main

//imports
import "fmt"

//package scope vars
type magicfish int

var x magicfish
var y int

//main func
func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
