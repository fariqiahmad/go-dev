package main

import "fmt"

func main() {
	var months = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	// var monthSlice = months[6:10]
	// fmt.Println(monthSlice)
	// fmt.Println(len(monthSlice))
	// fmt.Println(cap(monthSlice))

	// months[7] = "Juli-edited"
	// fmt.Println(monthSlice[1])

	// monthSlice[1] = "Juli-back"
	// fmt.Println(months[7])

	// var monthNew = months[10:11]
	// var monthNewAppend = append(monthNew, "Januari-new")
	// fmt.Println(monthNewAppend)
	// fmt.Println(monthNew)
	// fmt.Println(months)

	// monthNewAppend[0] = "November-new"
	// fmt.Println(monthNewAppend)
	// fmt.Println(monthNew)
	// fmt.Println(months)

	var monthSliceOut = months[11:]
	var monthSliceOutAppend = append(monthSliceOut, "New")
	fmt.Println(monthSliceOut)
	fmt.Println(monthSliceOutAppend)

	// Append will create new array for slice if more than capacity
	// and slice will refer to new array
	monthSliceOutAppend[0] = "Desember-edited"
	fmt.Println(monthSliceOut)
	fmt.Println(monthSliceOutAppend)
	fmt.Println(months)

	// Create new slice from new array
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ahmad"
	newSlice[1] = "F."
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	thisIsArray := [5]int{1, 2, 3, 4, 5} // or [...]int{1, 2, 3, 4, 5}
	thisIsSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}
