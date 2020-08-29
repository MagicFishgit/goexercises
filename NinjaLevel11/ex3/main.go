package main

import "fmt"

type customErr struct {
	additionalInfotoError string //long var name to make the learning clearer for myself.
}

func (custError customErr) Error() string {
	return fmt.Sprintf("Heres the error: %v", custError.additionalInfotoError)
}

func main() {

	c1 := customErr{
		additionalInfotoError: "Here's an error bud!",
	}

	foo(c1)

}

func foo(e error) {
	fmt.Println("this ran", e)
}
