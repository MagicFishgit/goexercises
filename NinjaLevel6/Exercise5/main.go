package main

import "math"

type circle struct {
	radius float64
}

type square struct {
	len float64
}

type shape interface {
	area() float64
}

func main() {

	//value of square
	squareval := square{
		len: 15,
	}

	circleval := circle{
		radius: 12.345,
	}

	//print info square
	info(squareval)
	info(circleval)

}

//Area method for square
func (sqr square) area() float64 {
	return sqr.len * sqr.len
}

//Area method for circle
func (circ circle) area() float64 {
	return math.Pi * (circ.radius * circ.radius)
}

//prints area from shape
func info(s shape) {
	println(int64(s.area()))
}
