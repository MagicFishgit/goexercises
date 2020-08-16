package main

import "fmt"

//type declaration
type person struct {
	first        string
	last         string
	iceCreamFlav []string
}

func main() {

	personMap := make(map[string]person)

	p1 := person{
		first:        "Rudi",
		last:         "Visagie",
		iceCreamFlav: []string{"Bubblegum", "Vanilla", "Bannana"},
	}

	personMap["Visagie"] = p1

	p2 := person{
		first:        "Sas",
		last:         "Swart",
		iceCreamFlav: []string{"Best bro in the universe", "Anxiety", "Sexy ass", "Oblivious to Emotions sometimes"},
	}

	personMap["Swart"] = p2

	fmt.Println(personMap["Visagie"].first, personMap["Visagie"].last)
	for i, v := range personMap["Visagie"].iceCreamFlav {
		fmt.Println(i, v)
	}

	fmt.Println(personMap["Swart"].first, personMap["Visagie"].last)
	for i, v := range personMap["Swart"].iceCreamFlav {
		fmt.Println(i, v)
	}
}

//I accidentally realised that I could do anonymous structs and assign them directly to the map thnx to my java objectliteral syntax understanding but I deverted that code to what the
//exercise question asked. Pretty impressed at myself lol.
