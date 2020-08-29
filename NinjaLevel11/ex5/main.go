package main

import "fmt"

func main() {
	//divide by x

	fmt.Println(Divide(100, 0))

}

//Divide testing go doc
func Divide(x, y float64) float64 {
	return (x / y)
}

//now im going to create an automated test and test dividing by 0
