package main

import "math"

type circle struct {
	radius float64
}

type square struct {
	len   float64
	width float64
}

type shape interface {
	area() float64
}

func main() {

	//value of square
	squareval := square{
		len:   10,
		width: 2,
	}

	circleval := circle{
		radius: 2,
	}

	//print info square
	info(squareval)
	info(circleval)

}

//Area method for square
func (sqr square) area() float64 {
	return sqr.len * sqr.width
}

//Area method for circle
func (circ circle) area() float64 {
	return math.Pi * math.Sqrt(circ.radius)
}

//prints area from shape
func info(s shape) {
	println(int64(s.area()))
}
