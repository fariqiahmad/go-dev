package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Ahmad F.",
		"address": "Jakarta",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["age"] = "25"
	fmt.Println(person["age"])

	fmt.Println(len(person))

	newMap := make(map[string]int)
	newMap["value1"] = 1
	newMap["value2"] = 2
	fmt.Println(newMap)
	delete(newMap, "value2")
	fmt.Println(newMap)
}
