package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result string = random().(string)
	fmt.Println(result)
	// menggunakan switch
	switch value := random().(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	}
}
