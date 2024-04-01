package main

import "fmt"

func main() {
	// Struktur slice {Pointer, Length, Capacity}
	// Pointer: penunjuk data pertama
	// Length: panjang slice
	// Capacity: kapasitas slice
	fruits := [...]string{
		"Mangga",
		"Pisang",
		"Jambu",
		"Nanas",
		"Semangka",
		"Anggur",
	}
	fmt.Printf("fruits: %v\n", fruits)
	slice1 := fruits[3:5]
	fmt.Println("Panjang slice1:", len(slice1))
	fmt.Println("Kapasitas slice1:", cap(slice1))
	fmt.Printf("slice1: %v\n", slice1)
	slice2 := fruits[:3]
	fmt.Printf("slice2: %v\n", slice2)
	slice3 := fruits[3:]
	fmt.Printf("slice3: %v\n", slice3)
	slice4 := fruits[:]
	fmt.Printf("slice4: %v\n", slice4)

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	fmt.Printf("days: %v\n", days)
	daysSlice1 := days[5:]
	daysSlice1[0] = "libur sabtu"
	daysSlice1[1] = "libur minggu"
	fmt.Printf("daysSlice1: %v\n", daysSlice1)
	fmt.Printf("days: %v\n", days)
	fmt.Println("\ndaysSlice1 masih mengikuti days jika masih dalam capasitas, jika lebih akan membuat array baru dan tidak merefer pada days lagi\n")
	daysSlice2 := append(daysSlice1, "libur nasional")
	daysSlice2[0] = "sabtu awal"
	fmt.Printf("daysSlice2: %v\n", daysSlice2)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "test 1"
	newSlice[1] = "test 2"
	fmt.Printf("newSlice: %v\n", newSlice)
	fmt.Println("Panjang newSlice:", len(newSlice))
	fmt.Println("Kapasitas newSlice:", cap(newSlice))
	newSlice2 := append(newSlice, "test 3")
	fmt.Printf("newSlice2: %v\n", newSlice2)
	fmt.Println("Panjang newSlice2:", len(newSlice2))
	fmt.Println("Kapasitas newSlice2:", cap(newSlice2))
	newSlice2[0] = "test 1 baru"
	fmt.Printf("newSlice: %v\n", newSlice)
	fmt.Printf("newSlice2: %v\n", newSlice2)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Printf("fromSlice: %v\n", fromSlice)
	fmt.Printf("toSlice: %v\n", toSlice)
}
