package main

import (
	"fmt"
)

type str string // custom type being created or also referred to as type aliases

func (text str) log() { // one can assign a method to a custom type which is not possible for a standard type
	fmt.Println(text)
}

func main() {
	var name str = "Ehis" // the custom type being used
	name.log()
}
