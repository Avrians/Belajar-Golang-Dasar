package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main34()  {
	var customer1 Customer
	customer1.Name = "John"
	customer1.Address = "LosAngeles"
	customer1.Age = 25
	fmt.Println(customer1)
	fmt.Println(customer1.Name)
	fmt.Println(customer1.Address)
	fmt.Println(customer1.Age)

	var customer2 Customer
	fmt.Println(customer2)

	// kita juga bisa bikin langusng

	joko :=  Customer{
		Name: "Joko",
		Address: "Bandung",
		Age: 30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println(budi)
	budi.sayHello("Agus")
	joko.sayHello("Budi")
}