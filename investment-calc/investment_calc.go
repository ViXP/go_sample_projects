package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64
	years, expectedReturnRate = 10, 5.5

	outputText("Set your investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Set your number of years: ")
	fmt.Scan(&years)

	outputText("Set your expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Future real value (inflation adjusted): %.2f", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(amount, rate, years float64) (fv, rfv float64) {
	fv = amount * math.Pow(1+rate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}
