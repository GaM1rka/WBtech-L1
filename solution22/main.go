package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2000000000000000)
	b := big.NewInt(52525525252)

	div := new(big.Int)
	div.Div(a, b)
	fmt.Println("Result of division 2 large numbers: ", div)

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Println("Result of adding 2 large numbers: ", sum)

	sub := new(big.Int)
	sub.Sub(a, b)
	fmt.Println("Result of substracting 2 large numbers: ", sub)

	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Println("Result of multiplication 2 large numbers: ", mul)
}
