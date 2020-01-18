package main

import "fmt"

func main() {
	x := 5
	y := 10

	// if-else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// switch
	color := "GREEN"
	switch color {
	case "red":
		fmt.Println("RED")
	case "blue":
		fmt.Println("RED")
	default:
		fmt.Println("UNKNOWN")
	}
}
