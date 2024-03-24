package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const USER_BALANCE_FILENAME = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(USER_BALANCE_FILENAME)
	if err != nil {
		return 0, errors.New("balance file not found")
	}
	balanceString := string(data)
	balance, err := strconv.ParseFloat(balanceString, 64)
	if err != nil {
		return 0, errors.New("balance not a valid number")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(USER_BALANCE_FILENAME, []byte(balanceString), 0644)
}

func main() {
	userBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for {
		var userChoice int
		showOptions()
		fmt.Print("Your choice: ")
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			fmt.Printf("Your current balance is %.2f\n", userBalance)
		} else if userChoice == 2 {
			var depositAmount float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. It has to be greater than 0.")
				continue
			}
			userBalance = userBalance + depositAmount
			writeBalanceToFile(userBalance)
			fmt.Printf("Balance updated! New amount: %.2f\n", userBalance)
		} else if userChoice == 3 {
			var withdrawAmount float64
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 || withdrawAmount > userBalance {
				fmt.Println("Invalid withdrawal amount")
				continue
			}
			userBalance = userBalance - withdrawAmount
			writeBalanceToFile(userBalance)
			fmt.Printf("Balance updated! New amount: %.2f\n", userBalance)
		} else if userChoice == 0 {
			return
		} else {
			fmt.Println("Invalid choice")
			break
		}
	}
}

func showOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("0. Exit")
}
