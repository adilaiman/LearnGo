package main

import "fmt"

// param followed by return type
func factorial(number uint64) uint64 {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
}

func sum(a, b int) (result int) {
	result = a + b
	return
}

func swapVals(a, b int) (result1, result2 int) {
	result1, result2 = b, a
	return
}

// variadic functions take a number of params
// Println is an example : Println(a, b, c ...)
func adder(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(sum(5, 3))
	fmt.Println(swapVals(5, 3))

	fmt.Println(adder(1, 2, 3, 4))

	// variadic functions can also be applied to slices
	numbers := []int{1, 2, 3, 4}
	fmt.Println(adder(numbers...))
}
