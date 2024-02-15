package main

import (
	"fmt"
)

func main() {
	age := 32
	fmt.Println("Age: ", age) // here a copy of age is what is created and passed into fmt.Println()

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int { // here another copy of age is what is created and passed into getAdultYears()
	return age - 18
}

// without the use of pointers, each time age is used a copy of age is created which can contribute immensely to large memory usage
