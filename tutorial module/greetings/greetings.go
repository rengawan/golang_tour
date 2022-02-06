package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello mengembalikan sebuah salam untuk nama seseorang.
func Hello(name string) (string, error) {
	// Jika nama kosong, kembalikan sebuah eror dengan pesan.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Buat sebuah pesan salam dengan format acak.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos mengembalikan sebuah map yang mengasosiasikan setiap nama orang
// dengan sebuah pesan salam.
func Hellos(names []string) (map[string]string, error) {
	// Sebuah map yang memetakan nama dengan pesan.
	messages := make(map[string]string)

	// Iterasi slice "names", panggil fungsi Hello untuk mendapatkan
	// sebuah pesan untuk setiap nama.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// Di dalam map, asosiasikan nama (kunci) dengan pesan (nilai).
		messages[name] = message
	}
	return messages, nil
}

// init men-set pengacak angka.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat mengembalikan sekumpulan pesan acak.  Pesan yang dikembalikan
// dipilih secara acak.
func randomFormat() string {
	// Slice dari sekumpulan format pesan.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Kembalikan salah satu format pesan secara acak.
	return formats[rand.Intn(len(formats))]
}
