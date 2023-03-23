package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers = []int{2, 3, 5, 7, 11, 13, 17, 19}
	var str = ""
	for i, number := range numbers {
		if i%2 == 1  {
			continue
		}
		if len(str) > 0 {
			str += "-"
		}
		str += strconv.Itoa(number)
	}
	fmt.Println(str)
}
