package main

import (
	"fmt"
)

func main() {
	age := 32 // regular variable

	agePointer := &age

	fmt.Println("Age: ", *agePointer)
	editAgeAdultYears(agePointer)
	fmt.Println("Age after mutation:", age)
}

func editAgeAdultYears(age *int) {
	*age = *age - 18
}
