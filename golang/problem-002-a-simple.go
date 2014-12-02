package main

import (
	"fmt"
)

func main() {
	ceiling := 4000000

	sum := sumEvenFibsUpTo(ceiling)

	fmt.Println("Sum of even fibs up to", ceiling, "is", sum)
}

func sumEvenFibsUpTo(ceiling int) int {
	sum := 0

	for _, i := range fibsUpTo(ceiling) {
		if i%2 == 0 {
			sum += i
		}
	}

	return sum
}

func fibsUpTo(ceiling int) []int {
	fibs := make([]int, 0)

	for i, j := 1, 1; i < ceiling; i, j = i+j, i {
		fibs = append(fibs, i)
	}

	return fibs
}
