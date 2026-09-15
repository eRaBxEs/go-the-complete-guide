package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	


	fmt.Print("Revenue amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses amount: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses

	earningsAfterTax := earningsBeforeTax - (earningsBeforeTax*taxRate/100)

	ratio := earningsBeforeTax/earningsAfterTax


	fmt.Println("Earnings before tax:", earningsBeforeTax)
	fmt.Println("Earnings after tax (profit):", earningsAfterTax)
	fmt.Println("Ratio of EBT/Profit:", ratio)
	

}