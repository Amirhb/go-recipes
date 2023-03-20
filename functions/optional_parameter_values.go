package main

import "fmt"

func main() {
	sayGoodbye1("")
	sayGoodbye1("see you")

	sayGoodbye2()
	sayGoodbye2("see you")
}

func sayGoodbye1(message string) {
	if len(message) == 0 {
		fmt.Println("Goodbye!")
	} else {
		fmt.Println(message)
	}
}

func sayGoodbye2(message ...string) {
	if len(message) == 0 {
		fmt.Println("Goodbye!")
	} else {
		fmt.Println(message[0])
	}
}
