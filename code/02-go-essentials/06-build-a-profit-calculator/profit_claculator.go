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
	earningsBeforeTax := 0.0
	profit := 0.0
	ratioEarningsToProfit := 0.0

	fmt.Print("Enter Revenue:")
	fmt.Scan(&revenue)
	fmt.Println()

	fmt.Print("Enter Expenses:")
	fmt.Scan(&expenses)
	fmt.Println()

	fmt.Print("Enter tax rate:")
	fmt.Scan(&taxRate)
	fmt.Println()

	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - (taxRate / 100))
	ratioEarningsToProfit = earningsBeforeTax / profit

	fmt.Println("Profit:", profit)
	fmt.Println("EBT:", earningsBeforeTax)
	fmt.Printf("ratio EBT/profit:%f/%f = %f", earningsBeforeTax, profit, ratioEarningsToProfit)

}
