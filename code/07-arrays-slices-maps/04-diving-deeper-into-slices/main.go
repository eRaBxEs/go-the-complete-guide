package main

import (
	"fmt"
)

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 2.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[1:] // a slice is more of a reference to an array
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))       // The length gives us the number of items in a slice or array
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // The capacity gives us the capacity of the underlying  for a slice

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // The capacity gives us the capacity of the underlying  for a slice
}
