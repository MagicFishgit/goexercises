package main

import "fmt"

func main() {
	const (
		_     = iota
		year1 = 2016 + iota
		year2 = 2016 + iota
		year3 = 2016 + iota
		year4 = 2016 + iota
	)

	fmt.Println(year1, year2, year3, year4)
}

//hmmmm interesting. But idk if i would really use it. We'll have to see in the future.
