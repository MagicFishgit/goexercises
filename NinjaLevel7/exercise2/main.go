package main

type person struct {
	first string
	last  string
}

func main() {

	x := person{
		first: "rudi",
		last:  "visagie",
	}

	println(x.first, x.last)
	changeme(&x)
	println(x.first, x.last)

}

func changeme(p *person) {

	(*p).first = "Hahahaha"

}
