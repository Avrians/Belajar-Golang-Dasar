package main

func main() {
	name := "Budi"

	switch name {
	case "Joko":
		println("Hello Joko")
	case "Budi":
		println("Hello Budi")
	default:
		println("Hi, Boleh kenalan?")
	}

	// Short Stement
	switch length := len(name); length > 5 {
	case true:
		println("Nama terlalu panjang")
	case false:
		println("Nama sudah benar")
	}

	// Switch tanpa kondisi
	switch {
	case len(name) > 5:
		println("Nama terlalu panjang")
	case len(name) < 5:
		println("Nama terlalu pendek")
	default:
		println("Nama sudah benar")
	}
}