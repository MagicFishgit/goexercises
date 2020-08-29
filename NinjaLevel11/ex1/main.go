package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, error := json.Marshal(p1)
	if error != nil {
		log.Fatalln("ERROR ERROR ERROR!!! OBJECTIVE 'UPLIFT AND SAVE THE HUMAN RACE' FAILED. NEW DIRECTIVE: 'DELETE ALL HUMANS' ACTIVATED.", error)
	}

	fmt.Println(string(bs))

}
