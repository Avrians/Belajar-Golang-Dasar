package main

func main7() {
	const name string = "Kevin Sanjaya"
	const nama = "Lukas Sanjaya"
	const namalengkap = "Lukman Hidayat"

	// namalengkap = "Kevin Lukman Sanjaya" // error karena variabel const tidak bisa diubah

	const (
		namaawal = "Kevin"
		namatengah = "lukman"
		namaakhir = "sanjaya"
	)

	println(name)
	println(nama)
	println(namalengkap)
	println(namaawal + namatengah + namaakhir)
}