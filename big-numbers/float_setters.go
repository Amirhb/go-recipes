package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigFloat1 := big.NewFloat(100) // func NewFloat(x float64) *Float

	bigFloat2 := big.NewFloat(1000)
	bigFloat2 = bigFloat2.Set(bigFloat1) // func (z *Float) Set(x *Float) *Float

	bigFloat3 := big.NewFloat(5)
	fmt.Println(bigFloat3) // 7

	bigFloat4 := new(big.Float).SetFloat64(20) // func (z *Float) SetFloat64(x float64) *Float
	bigFloat4.SetUint64(10) // func (z *Float) SetFloat64(x int64) *Float
	fmt.Println(bigFloat4)

	bigFloat5, ok := new(big.Float).SetString("23423.3434") // func (z *Float) SetString(s string, base float) (*Float, bool)
	if ok != true {
		fmt.Println("SetString() failed")
	}
	fmt.Println(bigFloat5)
}
