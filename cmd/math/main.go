package main

import (
	"fmt"

	"github.com/Team694/adder"
)

func main() {
	sum := adder.Add(1, 1)
	fmt.Println("1 + 1 =", sum)

	product := adder.Multiply(2, 2)
	fmt.Println("2 * 2 =", product)

	quotient := adder.Divide(2, 2)
	fmt.Println("2 / 2 =", quotient)
}
