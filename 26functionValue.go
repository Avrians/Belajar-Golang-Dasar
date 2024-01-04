package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main26()  {
	contoh := getGoodBye
	misal := getGoodBye
	fmt.Println(contoh("Rizky"))
	fmt.Println(misal("Rizky"))
}