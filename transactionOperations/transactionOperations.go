package transactionOperations

import (
	
	"fmt"
)



// //check if the user is already registered
// func isUserRegistered(userName string) bool {
// 	// if the userName is in the DB then check the passwords
// 	// isPasswordCorrect
// 	if (userName != "A") || (userName != "C") || (userName != "B") {
// 		return false
// 	}
// 	return true
// }

//check if the password is correct
// func isPasswordCorrect(password string) bool {
// 	if password != "hello" {
// 		return false
// 	}
// 	return true
// }

// deposit money
func DepositMoney(userName string, depositMoney float64) string {
	// Get current money from user account + deposit
	// set the money of account to total
	s := fmt.Sprintf("%f",depositMoney)
	var notification string = "User" + userName + "Account balance is:" + s + "."
	return notification
}

// Withdraw money
func WithdrawMoney(userName string, withdrawAmount float64) string {
	// Get current money from user account - withdraw
	// set the money of account to total
	s := fmt.Sprintf("%f",withdrawAmount)
	var notification string = "User" + userName + "Account balance is:" + s + "."
	return notification
}

// Transfer money
func TransferMoney(userNameA string, userNameB string, transferAmount float64) string {
	// Get current money from user A account
	// Get current money from user B account
	// Withdraw money from user A
	// withdrawMoney(userA, money)
	// Deposit money to user B
	// depositMoney(userB, money)
	// set the money of account to total
	s := fmt.Sprintf("%f",transferAmount)
	var notification string = "User" + userNameA + "Account balance is:" + s + "."
	return notification
}
