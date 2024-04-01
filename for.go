package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Counter", i)
	}

	// for range => (array, slice, dan map)
	cars := []string{
		"BMW",
		"Supra",
		"Jazz",
	}
	for index, car := range cars {
		fmt.Println("Mobil", (index + 1), "adalah", car)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println("index", (i + 1))
	}
}
