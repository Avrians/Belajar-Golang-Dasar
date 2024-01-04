package main

func main() {
	sayHelloTo("Eko", "Kurniawan")
	sayHelloTo("Budi", "Santoso")
}

func sayHelloTo(firstName string, lastName string) {
	println("Hello", firstName, lastName)
}