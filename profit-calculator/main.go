package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const calculationResultFile = "calculation.txt"

func main() {
	fmt.Println("***PROFITS CALCULATOR***")

	revenue, error := getUserValue("\nRevenue")

	if error != nil {
		panic(error)
	}

	expenses, error := getUserValue("Expenses")

	if error != nil {
		panic(error)
	}

	taxRate, error := getUserValue("Tax rate")

	if error != nil {
		panic(error)
	}

	earningBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)
	formattedEBT := fmt.Sprintf("\nEarning before tax: %.2f", earningBeforeTax)
	formattedProfit := fmt.Sprintf("\nProfit: %.2f", profit)
	formattedRatio := fmt.Sprintf("\nRatio: %.2f", ratio)

	if writeToFile([]string{formattedEBT, formattedProfit, formattedRatio}) != nil {
		panic("impossible to write to file!")
	}

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)
}

func writeToFile(formattedTexts []string) (writeError error) {
	joinedTexts := strings.Join(formattedTexts, "")
	writeError = os.WriteFile(calculationResultFile, []byte(joinedTexts), 0644)
	return
}

func getUserValue(caption string) (float64, error) {
	var userInput float64
	fmt.Print(caption, ": ")
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("the input can't be negative or zero")
	}

	return userInput, nil
}

func calculateValues(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - ebt*taxRate
	ratio = ebt / profit
	return
}
