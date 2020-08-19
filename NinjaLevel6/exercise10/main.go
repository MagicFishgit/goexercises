package main

func main() {
	a := "this is function level"
	func() {
		b := "this var is closured or closed off from the main func and can only run inside here"
		println(b)
	}()

	println(a)
	println("I cannot print b because I do not have access to it...")
}
