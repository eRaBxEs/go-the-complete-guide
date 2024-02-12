package main

import (
	"fmt"
)

func main() {
	// ask for revenue, expenses, tax-rate
	// calculate earnings before tax (EBT), earnings after tax (profit)
	// calcualte ratio EBT/profit
	// output profit, EBT and ratio

	earningsBeforeTax := 0.0
	profit := 0.0
	ratioEarningsToProfit := 0.0

	revenue := getUserInput("Revenue: ")

	expenses := getUserInput("Expenses: ")

	taxRate := getUserInput("Tax  Rate: ")

	earningsBeforeTax, profit, ratioEarningsToProfit = getFinancials(revenue, expenses, taxRate)

	fmt.Println("Profit:", profit)
	fmt.Println("EBT:", earningsBeforeTax)
	fmt.Printf("ratio EBT/profit:%f/%f = %f", earningsBeforeTax, profit, ratioEarningsToProfit)
}

func getUserInput(userInputText string) float64 {
	fmt.Println(userInputText)
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}

func getFinancials(revenue, expenses, taxRate float64) (earningsBeforeTax, profit, ratioEarningsToProfit float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - (taxRate / 100))
	ratioEarningsToProfit = earningsBeforeTax / profit
	return
}
