package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a * b
	fmt.Println(c)

	a /= 10
	fmt.Println(a)

	var d = 3
	var e = -d
	fmt.Println(+d)
	fmt.Println(e)
}
