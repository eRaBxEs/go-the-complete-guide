package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5 // if inflation rate was to be considered
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years = 10.0
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
