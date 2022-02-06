//nama paket/package
package main

//import paket/package lain
import (
	"fmt"
	"math/rand"
)

func main() {
	//pemanggilan fungsi paket
	fmt.Println("Bilangan kesukaan saya adalah", rand.Intn(10))
}
