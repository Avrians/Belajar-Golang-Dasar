package main

import "fmt"

func main13() {
	var names [3]string
	names[0] = "Kevin"
	names[1] = "Steven"
	names[2] = "Randy"

	println(names[0])
	println(names[1])
	println(names[2])

	// ketika diisi lebih dari 3, maka akan error
	//names[3] = "Budi" // error: index out of range [3] with length 3

	var values = [3]int{
		90,
		80,
		100,
	}
	fmt.Println(values)	
	
	var dataarray = [...]int{
		100,90,88,77,88,
	}

	fmt.Println(dataarray)
	fmt.Println(len(dataarray))
}