package main

import "fmt"

func main() {
	//divide by x

	fmt.Println(divide(100, 0))

}

func divide(x, y float64) float64 {
	return (x / y)
}

//now im going to create an automated test and test dividing by 0
