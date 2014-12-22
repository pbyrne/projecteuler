package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1001
	prime := nthPrime(n)

	fmt.Println("Prime number", n, "is", prime)
}

func nthPrime(n int) int {
	var result int
	var i = 1
	primes := make(chan int)
	go func() {
		feedPrimes(primes)
	}()

	for prime := range primes {
		if i == n {
			result = prime
			break
		}
		i += 1
	}

	return result
}

func feedPrimes(primes chan int) {
	for i := 2; ; i++ {
		if isPrime(i) {
			primes <- i
		}
	}
}

func isPrime(primeCandidate int) bool {
	if primeCandidate == 2 {
		return true
	}

	for i := 2; i <= squareRootOf(primeCandidate); i++ {
		if primeCandidate%i == 0 {
			return false
		}
	}

	return true
}

func squareRootOf(value int) int {
	return int(math.Ceil(math.Sqrt(float64(value))))
}
