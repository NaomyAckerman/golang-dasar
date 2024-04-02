package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func main() {
	person := Person{
		Name: "Novyan",
	}
	SayHello(person)
}
