package main

import "fmt"

//contoh variable dengan inisialisasi nilai awal pada level paket/package
var i, j int = 1, 2

func main() {

	//contoh variable dengan inisialisasi nilai awal pada level fungsi
	//type variable akan otomatis dengan nilai yang diberikan
	//seperti phyton dan java di isi dengan boolean true false

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
