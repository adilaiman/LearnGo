package main

import "fmt"

func main() {
	// Arrays is fixed and must have type
	// var fruitArr [2]string
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign, shorthand
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Slices, non fixed
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
