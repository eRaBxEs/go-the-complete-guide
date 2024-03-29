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

	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
}
