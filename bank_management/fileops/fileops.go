package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(fileName string, balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceString), 0644)
}
func GetBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
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
