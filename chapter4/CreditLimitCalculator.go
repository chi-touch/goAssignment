package main

import "fmt"

//var accountNumber, balanceAtBeginning, charges, credits, credit int
//var newBalance int
//
//func getNewBalance(balanceAtBeginning, charges, credits int) int {
//	newBalance := balanceAtBeginning + charges - credits
//	return newBalance
//}
//
//func checkIfNewBalanceExceedCredit(balanceAtBeginning, charges, credits int) string {
//	newBalance := balanceAtBeginning + charges - credits
//	if credit > newBalance {
//		return "Credit limit exceeded"
//	}
//	return "Credit limit is not exceeded"
//
//}

func main() {

	var accountNumber, balanceAtBeginning, charges, credits int
	var newBalance int

	fmt.Print("Please enter your account number: ")
	fmt.Scanln(&accountNumber)

	fmt.Print("Please enter your balance at beginning: ")
	fmt.Scanln(&balanceAtBeginning)

	fmt.Print("Enter total of all items charged by the customer this month: ")
	fmt.Scanln(&charges)

	fmt.Print("Enter allowed credit limit: ")
	fmt.Scanln(&credits)

	newBalance = balanceAtBeginning + charges - credits

	if newBalance > credits {
		fmt.Print("Credit limit exceeded")
	}

}
