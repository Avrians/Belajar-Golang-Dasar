package main

func main8() {
	// Konversi tipe data 1
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	println(nilai32)
	println(nilai64)
	println(nilai16)

	// Konversi tipe data 2
	var name = "Kevin Sanjaya"
	var e = name[0]
	var eString string = string(e)

	println(name)
	println(e)
	println(eString)
}