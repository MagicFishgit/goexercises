package main

//imports
import "fmt"

//vars package scope
var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
	fmt.Println("These values assigned by the compiler are called zero/default values and are assigned to variables when their type is assigned but no values are initialised.")
}
