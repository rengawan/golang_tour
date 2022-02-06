package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	//mengkonversi type integer ke float64
	var f float64 = math.Sqrt(float64(x*x + y*y))
	//mengkonversi float ke large integer
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
