package main

import (
	"fmt"
)

func main() {
	var value int
	factorMax := 20

	// brute-forces the thing since I don't know enough math to do otherwise
	for i := 1; ; i++ {
		if evenlyDivisbleUpTo(i, factorMax) {
			value = i
			break
		}
	}

	fmt.Println(value, "is divisible by all integers up to", factorMax)
}

func evenlyDivisbleUpTo(i int, factorMax int) bool {
	for j := 1; j <= factorMax; j++ {
		if i%j != 0 {
			break
		} else {
			if j == factorMax {
				return true
			}
		}
	}
	return false
}
