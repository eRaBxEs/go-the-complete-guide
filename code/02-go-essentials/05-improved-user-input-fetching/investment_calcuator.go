package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5 // if inflation rate was to be considered
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount:")
	fmt.Scan(&investmentAmount)

	fmt.Println()

	fmt.Print("Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Println()

	fmt.Print("Years:")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("Future Real Value: ", futureRealValue)
}
