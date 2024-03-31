package main

import "fmt"

func main() {
	type cString string

	var description cString = "Ini type alias dari string"
	fmt.Println(description)
}
