package main

import "fmt"

func main() {
	var str = ""
	var monitorInchSize = 24
	switch monitorInchSize {
		case 15:
			str = "too small"
		case 16, 17, 18:
			str = "good for the past decade"
		case 19, 20, 21, 22, 23:
			str = "for office work"
		case 24, 25, 26, 27:
			str = "great choice"
		default:
		str = ""
	}

	fmt.Println(str)
}
