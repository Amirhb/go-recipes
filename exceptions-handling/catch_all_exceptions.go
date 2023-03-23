package main

import (
	"fmt"
	"errors"
)

func main() {
	err := throwWhenNullOrEmpty(nil)

	if err != nil {
		fmt.Println(err)
	}
}

func throwWhenNullOrEmpty(arr []int) error {
	if arr == nil {
		return errors.New("Array is nil")
	}
	if len(arr) == 0 {
		return errors.New("Array is empty")
	}
	return nil
}
