package main

import (
	"fmt"
)

type Differ struct {
	ceiling int
}

func (d Differ) SquareOfSums() int {
	sum := 0

	for i := 1; i <= d.ceiling; i++ {
		sum += i
	}

	return sum * sum
}

func (d Differ) SumOfSquares() int {
	sum := 0

	for i := 1; i <= d.ceiling; i++ {
		sum += i * i
	}

	return sum
}

func main() {
	differ := Differ{100}

	fmt.Println("Square of sums up to", differ.ceiling, ": ", differ.SquareOfSums())
	fmt.Println("Sum of squares up to", differ.ceiling, ": ", differ.SumOfSquares())
	fmt.Println("Difference:", differ.SquareOfSums()-differ.SumOfSquares())
}
