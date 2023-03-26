package main

import "fmt"

type simpleException struct {
	errText string
}

func (e *simpleException) Error() string {
	return e.errText
}

func throwSimpleException() error {
	return &simpleException{"simple exception"}
}

func main() {
	err := throwSimpleException()

	if err != nil {
		fmt.Println(err.Error())
	}
}
