package main

import "fmt"

func main() {
	mappy := make(map[string][]string)

	mappy["James Bond"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	mappy["Rudi Visagie"] = []string{"Starts with a P and ends with and s", "Insecure.", "Good at solving other peoples emotional issues"}
	mappy["Sas Swart"] = []string{"Food", "Programming", "Ignoring people"}

	for i, v := range mappy {
		fmt.Printf("\nSubject: %v\n", i)
		for _, v := range v {
			fmt.Printf("Likes: %v\n", v)
		}
	}
}
