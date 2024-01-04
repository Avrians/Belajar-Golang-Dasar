package main

import "fmt"

func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main25() {
	total := sumAll(1, 2, 3, 4, 5)
	println(total)
	fmt.Println(sumAll(1, 2, 3, 4, 50))
	// fmt.Println(sumAll([]int{1, 2, 3, 4, 50}))])) jika membuat manual

	// kode program slice pada variadic function

	numbers := []int{1, 2, 3, 4, 5}
	total = sumAll(numbers...)
	fmt.Println(total)
	fmt.Println(sumAll(numbers...))
}