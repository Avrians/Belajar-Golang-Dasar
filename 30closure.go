package main

import "fmt"


func main30()  {
	counter := 0


	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	
	fmt.Println(counter)
}