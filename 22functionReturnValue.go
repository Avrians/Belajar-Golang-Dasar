package main

import "fmt"

func main22() {
	result := getHello("Eko")
	fmt.Println(result)
	fmt.Println(getHello("Budi"))
	fmt.Println(getHello("Joko"))
}

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}