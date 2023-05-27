package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAhmadF NoKTP = "978719275912"
	var marriedStatus Married = true

	fmt.Println(noKtpAhmadF, marriedStatus)
}
