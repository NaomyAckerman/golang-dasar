package main

import "fmt"

func main() {
	var (
		nilaiA      = 10
		nilaiB      = 20
		hasilJumlah = nilaiA + nilaiB
		hasilKali   = nilaiA * nilaiB
	)

	fmt.Println("Hasil dari ", nilaiA, "+", nilaiB, "=", hasilJumlah)
	fmt.Println("Hasil dari ", nilaiA, "*", nilaiB, "=", hasilKali)

	// Augmented assignments
	var scors = 9
	fmt.Println("Skor awal:", scors)
	scors += 3
	fmt.Println("Setelah skor awal ditambah 3:", scors)

	// Unary operator
	var increment = 9
	fmt.Println("Nilai awal:", increment)
	increment++
	fmt.Println("Nilai setelah increment:", increment)

	// Perbandingan
	var (
		nameA = "Vyan"
		nameB = "Vyan"
	)
	var result bool = nameA == nameB
	fmt.Println(nameA, "dan", nameB, "Merupakan nama yang", result)

	// Boolean
	var (
		nilaiUjian      = 90
		nilaiSikap      = 95
		lulusNilaiUjian = nilaiUjian >= 80
		lulusNilaiSikap = nilaiSikap >= 80
		lulus           = lulusNilaiSikap && lulusNilaiUjian
	)
	fmt.Println("Raport")
	fmt.Println("Nilai Ujian:", nilaiUjian, "(", lulusNilaiUjian, ")")
	fmt.Println("Nilai sikap:", nilaiSikap, "(", lulusNilaiSikap, ")")
	fmt.Println("Lulus:", lulus)
}
