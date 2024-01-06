package main

import "fmt"

func endApp()  {
	fmt.Println("End App")
	// Recover yang benar
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool)  {
	defer endApp()

	if error {
		panic("Upss Error")
	}

	// Recover yang salah
	// message := recover()
	// fmt.Println("terjadi panic", message)
}

func main()  {
	runApp(true)
	fmt.Println("OK")
}