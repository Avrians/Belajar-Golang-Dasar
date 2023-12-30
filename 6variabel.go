package main

import "fmt"

func main6(){
	var name string

	name = "Kevin Sanjaya"
	fmt.Println(name)

	var nama = "Lukas Sanjaya"
	fmt.Println(nama)

	namalengkap := "Lukman Hidayat" // hanya diawal untuk membuat variabel
	fmt.Println(namalengkap)

	var (
		namaawal = "Kevin"
		namatengah = "lukman"
		namaakhir = "sanjaya"
	)

	fmt.Println(namaawal + namatengah + namaakhir)
}