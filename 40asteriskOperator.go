package main

import "fmt"


// di off dulu karna di pake di 41operatorNew.go
// type Address struct {
// 	City, Province, Country string
// }


func main40() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer ke address1
	address2.City = "Bandung"
	fmt.Println(address1) // address1 berubah
	fmt.Println(address2) // address2 berubah

	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"} // ini tidak bisa
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // memindahkan data dari address2 ke address1
	fmt.Println(address1) // address1 berubah
	fmt.Println(address2) // address2 berubah
}