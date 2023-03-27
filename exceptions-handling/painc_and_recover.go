package main

import "fmt"

func panicAndDefer() {
	defer func() {
		panicText := recover()
		fmt.Println(panicText)
	}()

	panic("PANIC")
}

func main() {
	panicAndDefer()
}
