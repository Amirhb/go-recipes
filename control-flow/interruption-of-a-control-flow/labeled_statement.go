package main

import "fmt"

func main() {
	var firstMatchValue = -1
	var array1 = []int{1, 2, 3}
	var array2 = []int{2, 3, 4}

	FirstLoop:
	for _, value1 := range array1 {
		for _, value2 := range array2 {
			if value1 == value2 {
				firstMatchValue = value2
				break FirstLoop
			}
		}
	}

	fmt.Println(firstMatchValue)
}
