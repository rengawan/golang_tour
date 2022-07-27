package main

import "fmt"

func main() {
	var nilai32 int32 = 2000
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	//hati hati dalam konversi karena ketika overflow maka nilai akan berubah ke nilai  type yang paling kecil
	var nilai8 int8 = int8(nilai32)

	fmt.Println("Nilai 32 =", nilai32)
	fmt.Println("Nilai 64 =", nilai64)
	fmt.Println("Nilai 16 =", nilai16)
	fmt.Println("Nilai 8 =", nilai8)

	//perlakuan untuk string
	var nama = "Reno Gunawan"
	var e = nama[0]
	var eString string = string(e)

	println("Nama =", nama)
	println("eString =", eString)

}
