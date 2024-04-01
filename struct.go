package main

import "fmt"

type Car struct {
	Name, Type string
	Release    int16
}

// struct method
func (car Car) engineOn() {
	fmt.Println("Mesin mobil", car.Name, "Menyala")
}

func main() {
	var car Car
	fmt.Printf("car: %v\n", car)
	car.Name = "BWM"
	car.Type = "Car"
	car.Release = 2024
	fmt.Printf("car: %v\n", car)

	// struct literal
	car2 := Car{
		Name:    "Lambo",
		Type:    "Sport",
		Release: 2024,
	}
	fmt.Printf("car2: %v\n", car2)

	// struct method
	car2.engineOn()
}
