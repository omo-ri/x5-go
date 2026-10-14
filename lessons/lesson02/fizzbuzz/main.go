package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func main() {
	for i := range 30 {
		fmt.Println(FizzBuzz(i + 1))
	}
}
