package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	// any var or const declared in a function are scoped to that function
	var investmentAmount float64
	var expectedReturnRate float64
	years := 0.0


	// fmt.Print("Investment amount: ")
	outPutText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected return rate: ")
	outPutText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outPutText("Years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFuturevalues(investmentAmount, expectedReturnRate, years)

	// using Sprintf to return formatted string in variables
	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation):%.2f\n", futureRealValue)

	// fmt.Println("future value of investment:", futureValue)
	// fmt.Println("future real value after inflation:", futureRealValue)
	// fmt.Printf("Future value: %.2f\nFuture Value (adjusted for inflation):%.2f\n", futureValue, futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

}

func outPutText(text string) {
	fmt.Print(text)
}

func calculateFuturevalues(investmentAmount, expectedReturnRate, years float64) (fv, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv /  math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return
}