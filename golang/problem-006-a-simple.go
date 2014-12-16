package main

import (
	"fmt"
)

func main() {
	ceiling := 100
	sqOfSum := squareOfSums(ceiling)
	sumOfSq := sumOfSquares(ceiling)

	fmt.Println("Square of sums up to", ceiling, ": ", sqOfSum)
	fmt.Println("Sum of squares up to", ceiling, ": ", sumOfSq)
	fmt.Println("Difference:", sqOfSum-sumOfSq)
}

func squareOfSums(ceiling int) int {
	sum := 0

	for i := 0; i <= ceiling; i++ {
		sum += i
	}

	return sum * sum
}

func sumOfSquares(ceiling int) int {
	sum := 0

	for i := 0; i <= ceiling; i++ {
		sum += i * i
	}

	return sum
}
