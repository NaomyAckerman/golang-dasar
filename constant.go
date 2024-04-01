package main

import "fmt"

func main() {
	// Imutable
	const firstName = "Novyan"
	const lastName = "Dicky"

	const (
		job   = "Developer"
		hobby = "Coding"
	)

	fmt.Println("Nama depan:", firstName)
	fmt.Println("Nama belakang:", lastName)
	fmt.Println("Pekerjaan:", job)
	fmt.Println("Nama hobi:", hobby)
}
