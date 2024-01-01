package main

func main12() {
	var nilaiAkhir = 90
	var absensi = 81

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	println(lulus)
}