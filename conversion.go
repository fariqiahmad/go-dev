package main

import "fmt"

func main() {
	// consider ranges, value will converted if exceed of range
	var value16 int16 = 129
	var value8 int8 = int8(value16)
	var value32 int32 = int32(value16)

	fmt.Println(value16)
	fmt.Println(value8)
	fmt.Println(value32)

	var name = "Ahmad F."
	var fByte = name[6]
	var fChar = string(fByte)

	fmt.Println(name)
	fmt.Println(fByte)
	fmt.Println(fChar)
}
