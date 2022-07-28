//operator boolean
//Operator && dan ||

package main

import "fmt"

func main() {
	//Contoh Operator &&
	var ujian = 90
	var absensi = 80
	var lulusNilaiAkhir = ujian > 80
	var lulusAbsensi = absensi > 80
	lulus := lulusAbsensi && lulusNilaiAkhir
	fmt.Println("Status Lulus =", lulus)

	//COntoh Operator ||
	lulus = lulusAbsensi || lulusNilaiAkhir
	fmt.Println("Status Lulus =", lulus)
}
