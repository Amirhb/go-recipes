package main

import "fmt"

func main() {
	const (
		Platinum = iota
		Gold
		Silver
	)

	const (
		Mercury = iota + 1
		Venus
		Earth
	)

	var gold = Gold
	var venus = Venus

	fmt.Println(gold, venus)
}
