package main

func main9() {
	type NoKTP string

	var noKtpKevin NoKTP = "1234567890"
	var contohString = "2222222222"
	var contohKtpString = NoKTP(contohString)

	println(noKtpKevin)
	println(contohString)
	println(contohKtpString)
}