package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 4, 5, 6, 7}

	// loop through nums
	for index, num := range nums {
		fmt.Printf("index %d - %d\n", index, num)
	}

	// not using index
	for _, num := range nums {
		fmt.Printf("%d\n", num)
	}

	// summation
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Printf("Sum: %d\n", sum)

	// range with maps
	emails := map[string]string{"a": "a@a", "b": "b@b"}

	for key, val := range emails {
		fmt.Printf("%s: %s\n", key, val)
	}

	// just get key
	for key := range emails {
		fmt.Println("Key: " + key)
	}
}
