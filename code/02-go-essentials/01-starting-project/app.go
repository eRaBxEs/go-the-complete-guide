package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

// Go doesn't just know packages, but also modules
// Go modules can have many packages
// Every Go code needs to belong to a module atleast
// Also at the very least every go code must have a package main this is inorder to provide entry into the Go application as a whole
// When you have another file in your go module, it can also belong to the same main package which is indicated at the top of the code, 
// but there cannot be another main function as there can only be one main function even in the midst of many files belonging to the main package
