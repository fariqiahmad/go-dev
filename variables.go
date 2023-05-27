package main

import "fmt"

func main() {
	var name string

	name = "Ahmad F."
	fmt.Println(name)

	var company = "PT. ABC"
	fmt.Println(company)

	address := "Jakarta"
	fmt.Println(address)

	var (
		firstName = "Ahmad"
		lastName  = "F."
	)

	fmt.Println(firstName, lastName)
}
