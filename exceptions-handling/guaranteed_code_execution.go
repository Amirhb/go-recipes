package main

import (
	"errors"
	"fmt"
)

func errorIfTrue(param bool) (err error) {
	defer fmt.Println("defer")

	fmt.Println("start")
	if param {
		fmt.Println("error")
		return errors.New("test error")
	}

	return nil
}

func main() {
	errorIfTrue(true)
	// printed:
	// start
	// error
	// defer

	errorIfTrue(false)
	// printed:
	// start
	// defer
}
