package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file.")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balanve value.")
	}
	return value, nil
}

func writeFloatToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)

}

func main() {
	accountBalance, err := getFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
		// panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// checkBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Account Balance:", accountBalance)
		case 2:
			fmt.Print("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//return // use a naked return to stop execution
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawal amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return// use a naked return to stop execution
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}

	}

}
