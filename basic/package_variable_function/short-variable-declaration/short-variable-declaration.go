package main

import "fmt"

func main() {
	//deklarasi variable dengan inisialisasi namun masih dengan metode panjang
	var i, j int = 1, 2

	//deklarasi variable dengan metode pendek dan langsung memasukan nilai
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
