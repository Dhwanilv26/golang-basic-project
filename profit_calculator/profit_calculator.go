package main

import "fmt"

func main() {

	revenue := getUserInput("Revenue:")
	expenses := getUserInput("Expenses:")
	taxRate := getUserInput("Tax Rate:")

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(earningsBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit
	return earningsBeforeTax, profit, ratio

}

func getUserInput(infotext string) float64 {
	var userInput float64
	fmt.Println(infotext)
	fmt.Scan(&userInput)
	return userInput
}
