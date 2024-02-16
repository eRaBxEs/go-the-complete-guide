package main

import (
	"fmt"
)

func main() {
	age := 32

	agePointer := &age // to get the pointer of variable age use & in front of the value variable

	fmt.Println("Age: ", *agePointer) // Dereferencing: To get the value behind the pointer use * in front of the pointer variable

	getAdultYears(agePointer)
	fmt.Println("Age after mutation:", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
