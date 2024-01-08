package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // copy address1 ke address2

	address2.City = "Bandung"
	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2) // address2 berubah

	// pointer adalah reference ke variable lain
	address3 := &address1 // pointer ke address1
	address3.City = "Jakarta"
	fmt.Println(address1) // address1 berubah
	fmt.Println(address3) // address3 berubah

}