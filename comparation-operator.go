package main

import "fmt"

func main() {
	var name1 = "Ahmad"
	var name2 = "ahmad"
	var result1 = name1 == name2
	fmt.Println(result1)

	var val1 = 1
	var val2 = 5
	fmt.Println(val1 > val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 == val2)

	var var1 = 1
	var var2 = '1'
	var result2 = var1 == int(var2)
	fmt.Println(result2)
}
