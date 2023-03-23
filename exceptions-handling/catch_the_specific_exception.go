package main

import (
	"errors"
	"fmt"
)

func main() {
	err := throwIfNullOrEmpty(nil)

	if err.Error() == "Array is nil" {
		fmt.Println("Error: array is not specified")
	} else if err.Error() == "Array is empty" {
		fmt.Println("Error: array is empty")
	}
	// or:
	fmt.Println("Error:", err.Error())
}

func throwIfNullOrEmpty(arr []int) error {
	if arr == nil {
		return errors.New("Array is nil")
	}
	if len(arr) == 0 {
		return errors.New("Array is empty")
	}
	return nil
}
