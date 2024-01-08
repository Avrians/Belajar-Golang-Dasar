package helper

var version = "1.0.0" // tidak bisa diakses dari luar package
var Application = "Belajar Golang" // bisa diakses dari luar package
// nama harus diawali dengan huruf besar agar bisa diakses dari luar package

func SayHello(name string) string {
	return "Hello " + name
}

func sayGoodBye(name string) string {
	return "Good Bye " + name
}