package main

import (
  "fmt"
)

func main() {
  ceiling := 1000
  bases := []int{3,5}

  sum := sumMultiplesUpTo(ceiling, bases)

  fmt.Println("\n\nTotal sum is", sum)
}

func sumMultiplesUpTo(ceiling int, bases []int) int {
  sum := 0

  for i := 0; i < ceiling; i++ {
    sum += valueIfMultipleOfAny(i, bases)
  }

  return sum
}

func valueIfMultipleOfAny(value int, bases []int) int {
  for _, base := range bases {
    if value % base == 0 {
      return value
    }
  }

  return 0
}
