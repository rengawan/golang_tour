package main

import "fmt"

//dalam fungsi keluaran di berikan nama dengan variable x dan y
//x dan y langsung bisa di set di dalam function dan langsung di keluarkan
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
