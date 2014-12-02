package main

import (
	"fmt"
	"math"
)

func main() {
	value := 600851475143
	factor := largestPrimeFactorOf(value)

	fmt.Printf("The largest prime factor of %v is %v", value, factor)
}

func largestPrimeFactorOf(value int) int {
	var factor int

	for i := squareRootOf(value); i > 0; i-- {
		if isFactor(value, i) {
			if isPrime(i) {
				factor = i
				break
			}
		}
	}

	return factor
}

func squareRootOf(value int) int {
	return int(math.Sqrt(float64(value)))
}

func isFactor(value int, factorCandidate int) bool {
	return value%factorCandidate == 0
}

func isPrime(primeCandidate int) bool {
	for i := 2; i < squareRootOf(primeCandidate); i++ {
		if primeCandidate%i == 0 {
			return false
		}
	}

	return true
}
