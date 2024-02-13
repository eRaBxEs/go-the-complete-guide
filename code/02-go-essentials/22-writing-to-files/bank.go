package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {

	balanceText := fmt.Sprint(balance) // to be used to create the needed byte collections for the WriteFile function
	// os gives us a bunch of functions to interact with a file system
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = 1000.0
	fmt.Println("Welcome to Go bank!")

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				//return // use a naked return to stop execution
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New update:", accountBalance)
		case 3:
			fmt.Print("Your Withdrawal Amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue // use a naked return to stop execution
			}

			if withdrawalAmount > accountBalance { // check against withdrawing more than we have in account
				fmt.Println("Invalid amount. You can't withdraw more than you have")
				continue
			}

			accountBalance -= withdrawalAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New update:", accountBalance)
		default:
			fmt.Println("Good Bye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}

	}

}
