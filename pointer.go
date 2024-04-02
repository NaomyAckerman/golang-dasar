package main

import "fmt"

type Address struct {
	Country, Province, City string
}

func main() {
	/*
		pointer merupakan kemampuan
		reference ke lokasi data di memori yang sama
		tanpa menduplikasi data yang sudah ada
	*/
	// contoh pass by value
	address1 := Address{"Indonesia", "Jawa Timur", "Probolinggo"}
	address2 := address1 // pass by value
	address2.City = "Jember"
	fmt.Printf("address1: %v\n", address1) // tidak berubah
	fmt.Printf("address2: %v\n", address2) // berubah menjadi jember
	// contoh pass by reference (Pointer)
	var address3 Address = address1
	var address4 *Address = &address3 // (Pointer)
	address4.City = "Sidoarjo"
	fmt.Printf("address3: %v\n", address3) // berubah jika address4 berubah
	fmt.Printf("address4: %v\n", address4)

	// asterisk operator
	address5 := address1
	address6 := &address5
	address6.City = "Surabaya"
	fmt.Printf("address5: %v\n", address5)
	fmt.Printf("address6: %v\n", address6)
	// address6 = &Address{"Indonesia", "Jawa Timur", "Mojokerto"} // tidak merubah address5
	*address6 = Address{"Indonesia", "Jawa Barat", "Jakarta"} // merubah seluruh referencenya contohnya nilai address5
	fmt.Printf("address5: %v\n", address5)
	fmt.Printf("address6: %v\n", address6)

	// new pointer
	alamat1 := new(Address) // new = membuat pointer kosong
	alamat2 := alamat1
	alamat2.Country = "Indonesia"
	fmt.Printf("alamat1: %v\n", alamat1)
	fmt.Printf("alamat2: %v\n", alamat2)
}
