package main

import "fmt"

func Ups() any {
	return "Hello World"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
