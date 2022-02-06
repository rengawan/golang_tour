package main

import "fmt"

//function parameter dengan type sama
//cukup dengan mendeklarasikan di belakang parameter satu kali
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
