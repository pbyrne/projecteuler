package main

import (
	"fmt"
)

func main() {
	ceiling := 1000
	sum := 0

	for _, i := range multiplesUpTo(ceiling) {
		sum += i
	}

	fmt.Println("Total sum is", sum)
}

func multiplesUpTo(ceiling int) []int {
	values := make([]int, 0)

	for i := 0; i < ceiling; i++ {
		if i%3 == 0 || i%5 == 0 {
			values = append(values, i)
		}
	}

	return values
}
