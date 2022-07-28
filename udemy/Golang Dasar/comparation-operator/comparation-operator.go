//operator matematika golang

package main

import "fmt"

func main() {
	var name1 = "Reno"
	var name2 = "reno"
	var result bool = name1 == name2
	fmt.Println(result)

	var val1 = 100
	var val2 = 200

	fmt.Println(val1 < val2)
	fmt.Println(val1 > val2)
	fmt.Println(val1 == val2)
	fmt.Println(val1 != val2)

}
