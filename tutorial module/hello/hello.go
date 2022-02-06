package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properti dari Logger, termasuk prefiks dan opsi untuk mematikan
	// pencetakan waktu, sumber berkas, dan nomor baris.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Sebuah slice yang berisi nama-nama.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Panggil fungsi Hellos untuk mendapatkan pesan salam untuk setiap nama.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// Jika tidak ada eror, cetak map yang diterima dari ke layar.
	fmt.Println(messages)
}
