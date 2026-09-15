package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	years := 0.0


	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// using Sprintf to return formatted string in variables
	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation):%.2f\n", futureRealValue)

	// fmt.Println("future value of investment:", futureValue)
	// fmt.Println("future real value after inflation:", futureRealValue)
	// fmt.Printf("Future value: %.2f\nFuture Value (adjusted for inflation):%.2f\n", futureValue, futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

}