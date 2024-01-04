package main

import "fmt"

func main19(){
	// Break
	for i := 0; i < 5; i++ {
		if i == 5	{
			break
		}
		fmt.Println("Perulangan ke ", i)
	}

	// Continue
	for i := 0; i < 5; i++ {
		if i%2 == 0	{
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}