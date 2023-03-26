package main

import (
	"errors"
	"fmt"
)

func methodWithError() error {
	return errors.New("simple error")
}

func main() {
	err := methodWithError()
	fmt.Println(err)
}
