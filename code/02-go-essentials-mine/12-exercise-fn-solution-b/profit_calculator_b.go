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
	


	
	revenue = getConsoleFloatVal("Revenue amount: ")
	
	expenses = getConsoleFloatVal("Expenses amount: ")

	taxRate = getConsoleFloatVal("Tax rate: ")

	

	ebt, eat, ratio := getFinancials(revenue, expenses, taxRate)


	fmt.Println("Earnings before tax:", ebt)
	fmt.Println("Earnings after tax (profit):", eat)
	fmt.Printf("Ratio of EBT/Profit: %.3f\n", ratio)
	

}

func getConsoleFloatVal(inputFloatDescriptn string)(float64){
	fmt.Print(inputFloatDescriptn)
	var inputFloatVar float64
	fmt.Scan(&inputFloatVar)
	return inputFloatVar
}

func getFinancials(revenue, expenses, taxRate float64) (earningsBeforeTax, earningsAfterTax, ratio float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax - (earningsBeforeTax*taxRate/100)
	ratio = earningsBeforeTax/earningsAfterTax
	return
}



