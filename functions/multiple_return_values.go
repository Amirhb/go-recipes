package main

import "fmt"

func main() {
	ar := []int{2, 3, 5}
	first, last := getFirstLast(ar)

	fmt.Println("first =", first)
	fmt.Println("last =", last)
}

func getFirstLast(ar []int) (int, int) {
	var first, last = -1, -1
	if len(ar) > 0 {
		first = ar[0]
		last = ar[len(ar)-1]
	}
	return first, last
}
