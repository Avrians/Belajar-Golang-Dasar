package main

// func sayHelloWithFilter(name string, filter func(string) string)	{
// 	nameFiltered := filter(name)
// 	println("Hello", nameFiltered)
// }

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter)	{
	nameFiltered := filter(name)
	println("Hello", nameFiltered)
}

func spamFilter(name string) string	{
	if name == "Anjing"	{
		return "..."
	} else {
		return name
	}
}

func main()	{
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Budi", filter)
}