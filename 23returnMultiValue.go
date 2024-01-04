package main

import "fmt"

func main23() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// Kita bisa mengabaikan salah satu return value dengan menggunakan underscore (_)
	firstName2, _ := getFullName()
	fmt.Println(firstName2)
}

func getFullName() (string, string) {
	return "John", "Doe"
}