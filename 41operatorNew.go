package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = &Address{}
	var alamat2 *Address = new(Address) // jika data kosong, maka bisa menggunakan seperti ini
	// var alamat2 *Address = new(Address) // sama dengan alamat1

	alamat2.Country = "Indonesia"
	fmt.Println(alamat1) // alamat1 berubah
	fmt.Println(alamat2) // alamat2 berubah
}