package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("hellow")
	depositMoney("A", 300)
	withdrawMoney("B", 500)
	transferMoney("A", "B", 400)
	if isUserRegistered("A") {
		isPasswordCorrect("hello")
	}
}

//check if the user is already registered
func isUserRegistered(userName string) bool {
	// if the userName is in the DB then check the passwords
	// isPasswordCorrect
	if (userName != "A") || (userName != "C") || (userName != "B") {
		return false
	}
	return true
}

//check if the password is correct
func isPasswordCorrect(password string) bool {
	if password != "hello" {
		return false
	}
	return true
}

// deposit money
func depositMoney(userName string, depositMoney int) string {
	// Get current money from user account + deposit
	// set the money of account to total
	var notification string = "User" + userName + "Account balance is:" + strconv.Itoa(depositMoney) + "."
	return notification
}

// Withdraw money
func withdrawMoney(userName string, withdrawAmount int) string {
	// Get current money from user account - withdraw
	// set the money of account to total
	var notification string = "User" + userName + "Account balance is:" + strconv.Itoa(withdrawAmount) + "."
	return notification
}

// Transfer money
func transferMoney(userNameA string, userNameB string, transferAmount int) string {
	// Get current money from user A account
	// Get current money from user B account
	// Withdraw money from user A
	// withdrawMoney(userA, money)
	// Deposit money to user B
	// depositMoney(userB, money)
	// set the money of account to total
	var notification string = "User" + userNameA + "Account balance is:" + strconv.Itoa(transferAmount) + "."
	return notification
}
