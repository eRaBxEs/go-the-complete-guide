package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000.0

	fmt.Println("Welcome to Go bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	fmt.Println("Your Choice: ", choice)

	if choice == 1 {
		fmt.Println("Your account balance is", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0")
			return // use a naked return to stop execution
		}

		accountBalance += depositAmount
		fmt.Println("Balance updated! New update:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Your Withdrawal Amount:")
		var withdrawalAmount float64
		accountBalance -= withdrawalAmount
		fmt.Scan(&withdrawalAmount)
		fmt.Println("Balance updated! New update:", accountBalance)
	} else {
		fmt.Println("Good Bye!")
	}
}
