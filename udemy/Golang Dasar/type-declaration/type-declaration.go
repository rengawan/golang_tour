//golang mempunyai kemampuan untuk membuat alias dari type data yang sudah ada
package main

import "fmt"

func main() {
	//membuat type variable baru noKTP dan maritalStatus
	type noKTP string
	type maritalStatus bool

	var noKtp noKTP = "134646466546"
	var menikah maritalStatus = true

	fmt.Println("No KTP =", noKtp)
	fmt.Println("Status Menikah =", menikah)

}
