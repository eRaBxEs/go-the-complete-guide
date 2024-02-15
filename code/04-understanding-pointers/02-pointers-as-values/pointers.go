package main

import (
	"fmt"
)

func main() {
	age := 32

	var agePointer *int
	agePointer = &age // to get the pointer of variable age use & in front of the value variable

	fmt.Println("Age: ", *agePointer) // Dereferencing: To get the value behind the pointer use * in front of the pointer variable

	// adultYears := getAdultYears(age)
	// fmt.Println(adultYears)
}

func getAdultYears(age int) int { // here another copy of age is what is created and passed into getAdultYears()
	return age - 18
}
