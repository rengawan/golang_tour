//paket main default
package main

//import paket pendukung
import "fmt"

// func namafungsi (parameter type_parameter,parameter type_parameter) type_keluaran/output
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(68, 13))
}
