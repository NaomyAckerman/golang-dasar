package main

import "fmt"

func main() {
	// Imutable range after init range array
	// Tidak bisa menghapus dan tambah hanya bisa ubah
	var fruits [2]string
	fruits[0] = "Pisang"
	fruits[1] = "Mangga"

	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Printf("fruits: %v\n", fruits)

	var car = [...]string{
		"BMW",
		"Supra",
	}
	fmt.Println(car[0])
	fmt.Println(car[1])
	fmt.Printf("car: %v\n", car)
	fmt.Println("Total car:", len(car))
	car[1] = "Jazz"
	fmt.Printf("car terbaru: %v\n", car)
}
