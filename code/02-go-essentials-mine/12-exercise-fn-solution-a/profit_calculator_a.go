package main

import (
	"fmt"
)

func main() {
	// ask for revenue, expenses, tax-rate
	// calculate earnings before tax (EBT), earnings after tax (profit)
	// calcualte ratio EBT/profit
	// output profit, EBT and ratio
	
	var revenue float64
	var expenses float64
	var taxRate float64
	


	
	getConsoleFloatVal("Revenue amount: ", &revenue)
	
	getConsoleFloatVal("Expenses amount: ",&expenses)

	getConsoleFloatVal("Tax rate: ", &taxRate)

	

	ebt, eat, ratio := getFinancials(revenue, expenses, taxRate)


	fmt.Println("Earnings before tax:", ebt)
	fmt.Println("Earnings after tax (profit):", eat)
	fmt.Printf("Ratio of EBT/Profit: %.3f\n", ratio)
	

}

func getConsoleFloatVal(inputFloatDescriptn string, inputFloatVar *float64){
	fmt.Print(inputFloatDescriptn)
	fmt.Scan(inputFloatVar)
}

func getFinancials(revenue, expenses, taxRate float64) (earningsBeforeTax, earningsAfterTax, ratio float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax - (earningsBeforeTax*taxRate/100)
	ratio = earningsBeforeTax/earningsAfterTax
	return
}



