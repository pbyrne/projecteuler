package main

import (
	"fmt"
	"strings"
)

func main() {
	palindromes := make([]int, 0)

	for i := 999; i > 0; i-- {
		for j := 999; j > 0; j-- {
			product := i * j
			if isPalindromeNumber(product) {
				palindromes = append(palindromes, product)
			}
		}
	}

	product := palindromes[0]

	fmt.Printf("The largest palindrome product of two 3-digit numbers is %v", product)
}

func isPalindromeNumber(value int) bool {
	return string(value) == reversedString(value)
}

func reversedString(value int) string {
	stringy := string(value)
	length := len(stringy)
	result := make([]string, length)
	for i, c := range stringy {
		fmt.Printf("Reversing %v (length %v, index %v, fetching %v)", stringy, i, length-i)
		result[length-i] = string(c)
	}
	return strings.Join(result, "")
}
