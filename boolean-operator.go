package main

import "fmt"

func main() {
	var value = 80
	var attendance = 100

	var isPassValue bool = value > 75
	var isPassAttendance bool = attendance > 80

	var isPass bool = isPassValue && isPassAttendance
	fmt.Println(isPass)
}
