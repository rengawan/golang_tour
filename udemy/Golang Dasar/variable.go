package main

import "fmt"

func main() {
	var nama, alamat string
	var usia int8 = 38
	//golang akan otomatis inisialisasi type data apabila diberikan default
	var jmlanak = 1
	//deklarasi variable tapi var
	negara := "Indonesia"
	nama = "Reno Gunawan"
	alamat = "CLuster Sepatan CIty"
	fmt.Println("Nama = ", nama)
	fmt.Println("Usia = ", usia)
	fmt.Println("Alamat = ", alamat)
	fmt.Println("Negara = ", negara)
	fmt.Println("Mempunyai Anak = ", jmlanak)
}
