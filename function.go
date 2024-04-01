package main

import (
	"fmt"
	"strings"
)

func sayHello() {
	fmt.Println(`Hello World`)
}

// param
func sayHelloParam(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// return
func sayHelloReturn(name string) string {
	return "Hello " + name
}

// multiple return
func getFullname() (string, string) {
	return "Novyan", "Dicky"
}

// named return value
func getCompleteName() (firstName, lastName string) {
	firstName = "Novyan"
	lastName = "Dicky"
	return firstName, lastName
}

// variadic
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// function value
func getGoodBye(name string) string {
	return "Good bye " + name
}

// function as parameter
func sayHelloWithFilter(name string, filter func(string) string) string {
	return "Hello" + filter(name)
}
func spamFilter(s string) string {
	if strings.ToLower(s) == "anjing" {
		return "..."
	}
	return s
}

// funtion type declaration
type FilterSpam func(string) string

func sayHelloWithFilter2(name string, filter FilterSpam) string {
	return "Hello" + filter(name)
}
func spamFilter2(s string) string {
	if strings.ToLower(s) == "babi" {
		return "..."
	}
	return s
}

// anonymous function
type BlackList func(string) bool

func checkBlacklistUser(name string, blacklist BlackList) bool {
	if blacklist(name) {
		return true
	}
	return false
}

// recursive function
func faktorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * faktorialRecursive(value-1)
}

// defer
func logging() {
	fmt.Println("Selesai")
}
func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

// panic
func endApp() {
	fmt.Println("End Application 2")
}
func runApp(error bool) {
	defer endApp()
	fmt.Println("Run Application 2")
	if error {
		panic("Ups error")
	}
}

// panic
func endApp3() {
	fmt.Println("End Application 3")
	err := recover()
	if err != nil {
		fmt.Println("Terjadi error", err)
	}
}
func runApp3(error bool) {
	defer endApp3()
	fmt.Println("Run Application 3")
	if error {
		panic("Ups error")
	}
}

func main() {
	sayHello()

	// param
	sayHelloParam("Novyan", "Dicky")

	// return
	greeting := sayHelloReturn("Vyan")
	fmt.Println(greeting)

	// multiple return
	firstName, lastName := getFullname()
	firstName2, _ := getFullname()
	fmt.Println(firstName + " " + lastName)
	fmt.Println("Nama Depan:", firstName2)

	// named return value
	firstName3, lastName3 := getCompleteName()
	fmt.Println(firstName3, lastName3)

	// variadic
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10))
	numbers := []int{20, 20, 20, 20, 20}
	fmt.Println(sumAll(numbers...))

	// function value
	sayGoodbye := getGoodBye
	fmt.Println(sayGoodbye("Vyan"))

	// function as parameter
	sayHelloFilter := sayHelloWithFilter("Anjing", spamFilter)
	fmt.Println(sayHelloFilter)

	// function type declaration
	spamFilter2 := spamFilter2
	sayHelloFilter2 := sayHelloWithFilter2("Babi", spamFilter2)
	fmt.Println(sayHelloFilter2)

	// anonymous function
	isBlacklist := checkBlacklistUser("Vyan", func(s string) bool {
		if strings.ToLower(s) == "vyan" {
			return true
		}
		return false
	})
	fmt.Println("Apakah user Vyan termasuk dalam blacklist?", isBlacklist)

	// recursive function
	fmt.Println("Faktorial dari 10:", faktorialRecursive(10))

	// closure
	counter := 0
	fmt.Println("Konter awal", counter)
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	increment()
	increment()
	fmt.Println("Konter akhir", counter)

	// recover
	runApp3(true)

	// defer
	runApplication()

	// panic
	runApp(true)
}
