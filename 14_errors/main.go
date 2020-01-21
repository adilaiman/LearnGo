package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(number float64) (result float64, err error) {
	if number < 0 {
		return 0, errors.New("cannot find the square of a negative number")
	}
	return math.Sqrt(number), nil
}

func main() {
	fmt.Println(sqrt(-1))
	fmt.Println(sqrt(16))

	res, err := sqrt(64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
