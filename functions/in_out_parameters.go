package main

import "fmt"

func main() {
	var s1, s2 = "A", "B"
	swapString(&s1, &s2)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
}

func swapString(s1, s2 *string) {
	*s1, *s2 = *s2, *s1
}
