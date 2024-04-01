package main

import "fmt"

func main() {
	grade := "A"
	description := ""
	switch grade {
	case "A":
		description = "Grade A"
	case "B":
		description = "Grade B"
	case "C":
		description = "Grade C"
	default:
		description = "Tidak diketahui"
	}
	fmt.Println(description)

	// switch short statement
	name := "Novyan"
	switch length := len(name); length > 7 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama:", name)
	}

	length := len(name)
	switch {
	case length == 0:
		fmt.Println("Nama harus diisi")
	case length > 7:
		fmt.Println("Nama lebih dari 7 karakter")
	default:
		fmt.Println("Success Validation")
	}
}
