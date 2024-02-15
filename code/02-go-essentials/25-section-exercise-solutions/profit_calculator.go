package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals :
// 1. Validate users input
// show error message and exit if invalid input is provied
// No negative number
// No 0
// 2. Store calculated results into file

func main() {
	// ask for revenue, expenses, tax-rate
	// calculate earnings before tax (EBT), earnings after tax (profit)
	// calcualte ratio EBT/profit
	// output profit, EBT and ratio

	earningsBeforeTax := 0.0
	profit := 0.0
	ratioEarningsToProfit := 0.0

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax  Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, profit, ratioEarningsToProfit = getFinancials(revenue, expenses, taxRate)

	fmt.Println("Profit:", profit)
	fmt.Println("EBT:", earningsBeforeTax)
	fmt.Printf("ratio EBT/profit:%f/%f = %f", earningsBeforeTax, profit, ratioEarningsToProfit)

	err = writeResultToFile(earningsBeforeTax, profit, ratioEarningsToProfit)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
}

func getUserInput(userInputText string) (float64, error) {
	fmt.Print(userInputText)
	var userInput float64
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0.0, errors.New("Invalid input, should not be negative & must be greater than 0")
	}
	return userInput, nil
}

func getFinancials(revenue, expenses, taxRate float64) (earningsBeforeTax, profit, ratioEarningsToProfit float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - (taxRate / 100))
	ratioEarningsToProfit = earningsBeforeTax / profit
	return
}

func writeResultToFile(earningsBeforeTax, profit, ratioEarningsToProfit float64) error {
	resultText := fmt.Sprint("earningsBeforeTax:", earningsBeforeTax, "\n", "profit:", profit, "\n", "ratioEarningsToProfit:", ratioEarningsToProfit, "\n")
	err := os.WriteFile("financial-results.txt", []byte(resultText), 0644)
	if err != nil {
		fmt.Println("ERROR::", err)
		return err
	}
	return nil
}
