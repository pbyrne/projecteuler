package main

import (
	"fmt"
)

func main() {
	ceiling := 100
	sum := 0
	sqOfSum := 0
	sumOfSq := 0

	for i := 1; i <= ceiling; i++ {
		sumOfSq += i * i
		sum += i
	}
	sqOfSum = sum * sum

	fmt.Println("Square of sums up to", ceiling, ": ", sqOfSum)
	fmt.Println("Sum of squares up to", ceiling, ": ", sumOfSq)
	fmt.Println("Difference:", sqOfSum-sumOfSq)
}
