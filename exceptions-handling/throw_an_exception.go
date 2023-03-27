package main

import (
	"errors"
	"fmt"
)

type car struct {}

type seller struct {
	cars []car
}

func (s *seller) Sell() error {
	if len(s.cars) == 0 {
		return errors.New("No cars for sale")
	}

	return nil
}

func main() {
	var seller = seller{}

	if err := seller.Sell(); err != nil {
		fmt.Println(err)
	}
}
