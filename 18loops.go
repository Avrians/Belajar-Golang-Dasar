package main

import "fmt"


func main(){
	counter := 0

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	
	fmt.Println("Perulangan selesai")

	// for dengan statement
	for counter := 0; counter < 5; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}
	fmt.Println("Perulangan selesai")

	// for range
	names := []string{"Rizky", "Khapidsyah", "Go", "Lang"}

	// for biasa
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	// for range
	for index, name := range names {
		fmt.Println("Index ke ", index, "=", name)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}