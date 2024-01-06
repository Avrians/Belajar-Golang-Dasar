package main

import "fmt"

func endApp()  {
	fmt.Println("End App")
}

func runApp(error bool)  {
	defer endApp()

	if error {
		panic("Upss Error")
	}
}

func main()  {
	runApp(true)
}