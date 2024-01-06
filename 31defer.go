package main

import "fmt"

func logging()  {
	fmt.Println("Selesai memanggil function")
}

func runApplication()  {
	defer logging()
	fmt.Println("Run Application")

	// jika ini terjadi eror
	// logging()  // ini tidak akan dipanggil
	// diganti dengan yang diawal function
}

func main()  {
	// Defer
	runApplication()
}