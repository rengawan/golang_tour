//operator matematika golang

package main

func main() {
	var a = 30 + 10
	var b = a - 20
	var c = b*3 + 1
	var d = c % 2

	println(a)
	println(b)
	println(c)
	println(d)

	//augmented assignment
	var i = 10
	i += 30
	println(i)
	i -= 20
	println(i)

	//unari
	var j = 10
	j++
	println(j)
	j--
	println(j)
}
