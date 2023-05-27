package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Ahmad"
	name[1] = "F."

	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{0, 1, 3}
	var emptyValues [100]int
	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println((len(emptyValues)))
}
