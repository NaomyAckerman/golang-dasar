package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Novyan",
		"job":  "Programmer",
	}
	person["hobby"] = "Koding"
	fmt.Printf("person: %v\n", person)

	books := make(map[string]string)
	books["title"] = "Buku sakti"
	books["author"] = "Vyan"
	books["genre"] = "Edukasi"
	fmt.Printf("books: %v\n", books)
	delete(books, "genre")
	fmt.Println("Panjang data books", len(books))
	fmt.Printf("books: %v\n", books)
}
