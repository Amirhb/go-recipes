package main

import "fmt"

func main() {
	var n = -42
	// Go doesn't have the ternary operator
	var classify = (map[bool]string{true: "positive", false: "negative"})[n > 0]
	fmt.Println("classify:", classify)
}
