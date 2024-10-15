package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 30) // 2^30
	b := big.NewInt(1 << 20) // 2^20

	sum := new(big.Int)
	sub := new(big.Int)
	mult := new(big.Int)
	div := new(big.Int)

	sum.Add(a, b)
	sub.Sub(a, b)
	mult.Mul(a, b)
	div.Div(a, b)

	fmt.Printf("a: %s\n", a.String())
	fmt.Printf("b: %s\n", b.String())
	fmt.Printf("a + b: %s\n", sum.String())
	fmt.Printf("a - b: %s\n", sub.String())
	fmt.Printf("a * b: %s\n", mult.String())
	fmt.Printf("a / b: %s\n", div.String())
}
