package main

import "fmt"

func main() {
	a := 5
	// assign b to the pointer of a (mem address)
	b := &a
	fmt.Println(a, b)

	// *int - pointer int
	fmt.Printf("%T\n", b)

	// read val at mem address
	fmt.Println(*b)

	// change val at mem address
	*b = 10
	fmt.Println(a)
}
