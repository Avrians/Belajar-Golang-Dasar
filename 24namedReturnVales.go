package main

import "fmt"

func main24() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

func getCompleteName() (firstName, middleName, lastName string) { // jika name parameternya sama semua bisa dituliskan salah satu
	firstName = "Joko"
	middleName = "Eko"
	lastName = "Budi"

	return firstName, middleName, lastName
}