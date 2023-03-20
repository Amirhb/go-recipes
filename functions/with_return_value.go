package main

import "fmt"

func main() {
	var sum = getSum(5, 3)
	fmt.Println("sum =", sum)
}

func getSum(n1, n2 int) int {
	return n1 + n2
}
