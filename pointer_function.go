package main

import "fmt"

type Address struct {
	Country, Province, City string
}

func ChangeContryToIndonesia(address Address) {
	address.Country = "Indonesia"
}
func ChangeContryToIndonesiaPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address_pf := Address{}
	ChangeContryToIndonesia(address_pf)
	fmt.Printf("address_pf: %v\n", address_pf) // address_pf tidak berubah (bukan pointer)

	address_pf2 := Address{}
	ChangeContryToIndonesiaPointer(&address_pf2) // address_pf2 berubah
	fmt.Printf("address_pf2: %v\n", address_pf2)
}
