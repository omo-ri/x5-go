package main

import "fmt"

func Stats(money []int) (lo, hi, sum int, avg float64, ok bool) {
	if len(money) == 0 {
		return
	}
	ok = true
	lo, hi = money[0], money[0]
	for _, value := range money {
		if value < lo {
			lo = value
		}
		if value > hi {
			hi = value
		}
		sum += value
	}
	avg = float64(sum) / float64(len(money))
	return
}

func main() {
	money := []int{12000, 7500, 31000, 4500}
	lo, hi, sum, avg, ok := Stats(money)
	if ok {
		fmt.Printf("min=%d max=%d sum=%d avg=%.2f\n", lo, hi, sum, avg)
	} else {
		fmt.Println("Slice is empty")
	}
}
