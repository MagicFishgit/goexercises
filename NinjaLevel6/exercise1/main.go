package main

func main() {

	a, b := bar()

	println(foo(), a, b)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "a string"
}
