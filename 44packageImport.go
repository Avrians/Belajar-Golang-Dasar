package main

import "belajar-golang-dasar/helper"
import "fmt"

func main()  {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses karena nama variabel diawali huruf kecil
	// fmt.Println(helper.sayGoodBye("Eko") // tidak bisa diakses karena nama fungsi diawali huruf kecil
}