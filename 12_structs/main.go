package main

import "fmt"

// define person struct
type Person struct {
	firstName string
	lastName  string
	age       int
}

// greeting method (value receiver) "getter"
// p similar to "this"
func (p Person) greet() string {
	return "Hello from " + p.firstName + "."
}

// change age method (pointer receiver) "setter"
func (p *Person) changeAge(newAge int) {
	p.age = newAge
}

func main() {
	// init person using struct
	person1 := Person{firstName: "foo", lastName: "bar", age: 40}

	// shorthand
	person2 := Person{"fizz", "buzz", 30}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.firstName)
	fmt.Println(person2.age)

	fmt.Println(person1.greet())

	fmt.Println(person1.age)
	person1.changeAge(25)
	fmt.Println(person1.age)

	// pointers
	person1Ptr := &person1
	fmt.Println(person1Ptr.firstName)
	person1Ptr.firstName = "fizz"
	fmt.Println(person1.firstName)
}
