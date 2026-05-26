package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Revenue:")

	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses:")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate:") // here the latest err is used, so we need seperate checks for all invalid conditions
	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(earningsBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)

	storeResults(earningsBeforeTax, profit, ratio)

}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("ebt: %.1f\nProfit:%.1f\nRatio:%.1f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit
	return earningsBeforeTax, profit, ratio

}

func getUserInput(infotext string) (float64, error) {
	var userInput float64
	fmt.Println(infotext)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number.")
	}
	return userInput, nil
}
