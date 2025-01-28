package main

import (
	"fmt"
)

// struct is a collection of fields, is a way to group data
type Person struct {
	Name       string
	Age        int
	Profession string
}

// we can use a receiver to call the function
// is common to use the name of the struct as receiver
func (p Person) ListNameAndAge() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
	Person1 := Person{
		Name:       "John",
		Age:        30,
		Profession: "Engineer",
	}

	Person2 := Person{
		Name:       "Jane",
		Age:        25,
		Profession: "Doctor",
	}

	//calling the function with the receiver
	println(Person1.ListNameAndAge())
	println(Person2.ListNameAndAge())
}
