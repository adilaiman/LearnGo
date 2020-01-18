package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int int8 int16 int32 int64 (-n - n)
	// uint uint8 uint16 uint64 uintptr (0 - n)
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var
	var name = "foobar"
	// same as: var name string = "foobar"
	var age uint = 29 // vars have to be used when created
	const isEmployed = true
	// isEmployed = false, ERROR: const cannot be reassigned

	// cannot do := outside of a function
	// shorthand, same as var
	from := "USA"

	fmt.Println(name, age, from)
	fmt.Printf("type: %T", isEmployed) // get type and string formatting
}
