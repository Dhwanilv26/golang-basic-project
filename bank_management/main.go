package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	// even if the file doesn't exist, the program doesnt crash at all, the accountbalance will be assigned 0 (balance -> empty byte slice) .. ERRORS DONT CRASH THE COLLECTION

	if err != nil {
		return 1000, errors.New("failed to read file") // return default balance
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceString), 0644)
}

func main() {

	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		panic("there is no reason to continue")
	}

	fmt.Println("Welcome to Go Bank:")

	for {
		fmt.Println("What do you want to do?")

		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Println("Your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("your balance is:", accountBalance)
		case 2:
			fmt.Println("enter how much you want to deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("invalid amount, must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("balance updated, new balance:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("enter how much you want to withdraw:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("invalid amount, must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("withdraw amount cant be more than the account balance")
				return
			}
			accountBalance -= withdrawAmount
			fmt.Println("balance updated, new balance:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("goodbye!")
			return
		}
		fmt.Println("thanks for visiting our bank")
	}
}
