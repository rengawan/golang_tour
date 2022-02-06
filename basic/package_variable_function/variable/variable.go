package main

import "fmt"

//contoh variable dalam golang
//contoh variable di tulis pada tingkat paket/package yang bisa di akses secara global pada fungsi dibawahnya
var c, python, java bool

func main() {
	//contoh variable pada tingkat fungsi
	var i int
	fmt.Println(i, c, python, java)
}
