package main

import (
	"errors"
	"fmt"
)

var ErrNegativeAmount = errors.New("negative amount")

const (
	SmallCategory   = "small"
	RegularCategory = "regular"
	LargeCategory   = "large"
)

func Category(kopeyk int) (string, error) {
	switch {
	case kopeyk < 0:
		return "", ErrNegativeAmount
	case kopeyk < 10_000:
		return SmallCategory, nil
	case kopeyk < 50_000:
		return RegularCategory, nil
	default:
		return LargeCategory, nil
	}
}

func main() {
	money := []int{12000, 7500, 31000, 4500, -100, 100000, 25500}
	for _, v := range money {
		if cat, err := Category(v); err == nil {
			fmt.Println(v, cat)
		} else {
			fmt.Println("Error: ", err.Error())
			break
		}
	}
}
