package main

func main() {
	whatfish := magicfish()

	whatfish()
}

func magicfish() func() {

	return func() {
		println("I am a magical fish!")
	}

}
