package main

import "fmt"

func main() {
	// define map
	emails := make(map[string]string)

	// assign key values
	emails["b"] = "b@b.com"
	emails["a"] = "a@a.com"
	emails["c"] = "c@c.com"

	fmt.Println(emails)
	fmt.Println(emails["a"])
	fmt.Println(len(emails))

	// delete from map
	delete(emails, "b")
	fmt.Println(emails)

	// declare and define kv
	emails2 := map[string]string{"b": "b@b", "c": "c@c"}
	fmt.Println(emails2)
}
