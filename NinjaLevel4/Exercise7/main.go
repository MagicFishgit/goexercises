package main

import "fmt"

func main() {
	xCountries := make([]string, 50, 60)
	fmt.Println(len(xCountries))
	fmt.Println(cap(xCountries))

	//The question was asked wrong and does not include all the detail expected of the answer.
	//populate the slice

	for i := 0; i < len(xCountries); i++ {
		xCountries[i] = fmt.Sprintf("%d", i)
	}

	fmt.Println(xCountries)
	fmt.Println(len(xCountries))
	fmt.Println(cap(xCountries))

	xCountries = append(xCountries, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(xCountries)
	fmt.Println(len(xCountries))
	fmt.Println(cap(xCountries))

	for i := 0; i < len(xCountries); i++ {
		fmt.Println(i, xCountries[i])
	}

}
