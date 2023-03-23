package main

import "fmt"

func main() {
	var firstNumber = "1"
	var numberList = ""
	switch firstNumber {
		case "1":
			numberList = "1"
			fallthrough
		case "2":
			numberList = "-2"
			fallthrough
		case "3":
			numberList = "-3"
	}

	fmt.Println(numberList)
}

