package main

import "fmt"

func main() {
	nilaiUjian := 90
	if nilaiUjian >= 90 {
		fmt.Println("Grade A")
	} else if nilaiUjian >= 80 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	// if short statement
	phoneNumber := "081234567890"
	if length := len(phoneNumber); length <= 12 && length >= 9 {
		fmt.Println("No Telepon valid", phoneNumber)
	} else {
		fmt.Println("No Telepon tidak valid")
	}
}
