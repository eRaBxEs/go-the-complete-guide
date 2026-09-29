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

	// Imagine that we wanted to append balance one after the other to the file at each time
	writeAppendBalanceToFile("ebt", ebt)
	writeAppendBalanceToFile("profit", eat)
	writeAppendBalanceToFile("ratio", ratio)

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

func writeAppendBalanceToFile(amountDescription string, amount float64) {
	fileName := "financials.txt" // file name

	// create file if it does not exist
	_, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	// read the file next
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	amountText := fmt.Sprintf("%.2f", amount)
	amountDescriptionText := fmt.Sprint(amountDescription)
	finalData := fmt.Sprintln(amountDescriptionText, ":", amountText)

	dataToAppend := []byte(finalData)

	data = append(data, dataToAppend...)
	os.WriteFile(fileName, data, 0644)

}
