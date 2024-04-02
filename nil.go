package main

import "fmt"

/*
Nil hanya bisa pada type:
map, interface, function, slice, pointer, dan channel
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	name := NewMap("Hello Novyan")
	if name == nil {
		fmt.Println("Data null")
	} else {
		fmt.Printf("value: %v\n", name)
	}
}
