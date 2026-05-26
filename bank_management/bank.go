package main

import (
	"fmt"
	"example.com/bank-management/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, err := fileops.GetBalanceFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		panic("there is no reason to continue")
	}

	fmt.Println("Welcome to Go Bank:")

	for {
		fmt.Println("What do you want to do?")

		presentOptions()

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
			fileops.WriteBalanceToFile(accountBalanceFile, accountBalance)
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
			fileops.WriteBalanceToFile(accountBalanceFile, accountBalance)
		default:
			fmt.Println("goodbye!")
			return
		}
		fmt.Println("thanks for visiting our bank")
	}
}
