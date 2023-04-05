package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigInt1 := big.NewInt(100) // func NewInt(x int64) *Int

	bigInt2 := big.NewInt(1000)
	bigInt2 = bigInt2.Set(bigInt1) // func (z *Int) Set(x *Int) *Int

	bigInt3 := big.NewInt(5)
	bigInt3.SetBit(bigInt3, 1, 1) // SetBit sets z to x, with x's i'th bit set to b (0 or 1). That is, if b is 1 SetBit sets z = x | (1 << i); if b is 0 SetBit sets z = x &^ (1 << i). If b is not 0 or 1, SetBit will panic.
	fmt.Println(bigInt3) // 7

	bigInt4 := new(big.Int).SetInt64(20) // func (z *Int) SetInt64(x int64) *Int
	bigInt4.SetUint64(10) // func (z *Int) SetInt64(x int64) *Int
	fmt.Println(bigInt4)

	bigInt5, ok := new(big.Int).SetString("0000000101", 2) // func (z *Int) SetString(s string, base int) (*Int, bool)
	if ok != true {
		fmt.Println("SetString() failed")
	}
	fmt.Println(bigInt5)
}
