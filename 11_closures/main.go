package main

import "fmt"

// closure - returns anon function

// closure ex 1
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// closure ex 2
func returnAnonFunc() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}

func main() {

	// anon function
	func(msg string) {
		fmt.Println(msg)
	}("Hello World!")

	// closure ex 1
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

	// closure ex 2
	printMsg := returnAnonFunc()
	printMsg("Hello World2!")
}
