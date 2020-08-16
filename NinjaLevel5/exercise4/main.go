package main

import "fmt"

func main() {

	fish := struct {
		animal    string
		isMagical bool
		powerStr  map[string]int
	}{
		animal:    "Fish",
		isMagical: true,
		powerStr: map[string]int{
			"Vanish":               9,
			"teleport":             9,
			"turn wine into water": 1,
		},
	}

	fmt.Println(fish)
}
