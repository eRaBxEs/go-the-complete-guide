package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// Goals
	// 1) Validate user inout
	//  => Show error message and exit if invalid input is provided
	// - No negative numbers
	// - Not 0
	// 2) Store calculated results into file

	revenue, err := getConsoleFloatVal("Revenue amount: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expenses, err := getConsoleFloatVal("Expenses amount: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	taxRate, err := getConsoleFloatVal("Tax rate: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ebt, eat, ratio := getFinancials(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax:", ebt)
	fmt.Println("Earnings after tax (profit):", eat)
	fmt.Printf("Ratio of EBT/Profit: %.3f\n", ratio)

	storeResultsToToFile(ebt, eat, ratio)

}

func getConsoleFloatVal(inputFloatDescriptn string) (float64, error) {
	fmt.Print(inputFloatDescriptn)
	var inputFloatVar float64
	fmt.Scan(&inputFloatVar)

	if inputFloatVar <= 0 {
		return 0, errors.New("negative values or zero are not allowed")
	}
	return inputFloatVar, nil
}

func getFinancials(revenue, expenses, taxRate float64) (earningsBeforeTax, earningsAfterTax, ratio float64) {
	earningsBeforeTax = revenue - expenses

	earningsAfterTax = earningsBeforeTax - (earningsBeforeTax * taxRate / 100)

	ratio = earningsBeforeTax / earningsAfterTax

	return
}

func storeResultsToToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
