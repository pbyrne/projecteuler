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
	fibs := fibGenerator(ceiling)

	for i := range fibs {
		if i%2 == 0 {
			sum += i
		}
	}

	return sum
}

func fibGenerator(ceiling int) chan int {
	fibs := make(chan int)

	go func() {
		for i, j := 1, 1; i < ceiling; i, j = i+j, i {
			fibs <- i
		}
		close(fibs)
	}()

	return fibs
}
