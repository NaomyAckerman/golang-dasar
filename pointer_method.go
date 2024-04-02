package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	person := Man{"Novyan"}
	person.Married()
	fmt.Printf("person: %v\n", person.Name)
}
