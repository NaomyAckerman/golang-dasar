package main

import "fmt"

func main() {
	// jika melebihi kapasitas akan direset ke nilai minimum
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Novyan"
	var firstChar = name[0]
	fmt.Println("Karakter pertama dari \"Novyan\":", "String:", string(firstChar), "Byte:", firstChar)
}
