package main

import "fmt"

// Factorial menggunakan Loop
func factorialLoop(n int) int {
	var result = 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

// Factorial menggunakan recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main()  {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}