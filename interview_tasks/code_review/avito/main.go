package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func changeName(person *Person) {
	*person = Person{
		Name: "James",
	}
}

func task1() {
	person := Person{
		Name: "Alice",
	}

	fmt.Println(person.Name) // Alice
	changeName(&person)
	fmt.Println(person.Name) // James
}
