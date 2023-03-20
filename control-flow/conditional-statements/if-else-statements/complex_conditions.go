package main

import "fmt"

func main() {
	var A, B, C = 3, 5, 7
	if C >= A && C >= B {
		fmt.Println("nothing is larger than C")
	}
	if !(A >= B || A >= C) {
		fmt.Println("A is the smallest")
	}
}
