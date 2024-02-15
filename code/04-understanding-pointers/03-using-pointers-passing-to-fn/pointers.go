package main

import (
	"fmt"
)

func main() {
	age := 32

	agePointer := &age // to get the pointer of variable age use & in front of the value variable

	fmt.Println("Age: ", *agePointer) // Dereferencing: To get the value behind the pointer use * in front of the pointer variable

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int { // here a copy of age is not what is created but rather the referenced value
	return *age - 18
}
