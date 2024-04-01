package main

import "fmt"

func main() {
	var name string
	var nickname = "vyan"
	job := "Developer"

	// Multiple variable
	var (
		firstName = "Novyan"
		lastName  = "Dicky"
	)

	name = "Novyan Dicky"
	fmt.Println("Nama:", name)
	fmt.Println("Nama panggilan:", nickname)
	fmt.Println("Pekerjaan:", job)
	fmt.Println("Nama Depan:", firstName)
	fmt.Println("Nama Belakang:", lastName)
}
