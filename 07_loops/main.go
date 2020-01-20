package main

import "fmt"

func main() {
	// long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number: %d\n", i)
	}

	// fizzbuzz
	// every multiple of 3 output fizz
	// every multiple of 5 output buzz
	// every multiple of both output fizzbuzz
	// up to range 100
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else {
			fmt.Println(i)
		}
	}
}
