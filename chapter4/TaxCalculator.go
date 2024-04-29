package main

import "fmt"

func main() {

	var totalTax, taxRate, balanceOfEarning, taxBalance, balance, tax, earning float64
	var name string

	for count := 0; count <= 3; count++ {
		fmt.Println("Enter your name: ")
		fmt.Scanln(name)

		fmt.Println("Enter your earning:")
		fmt.Scanln(earning)

		if earning <= 30000 || earning > 30000 {
			taxRate = (earning * 15) / 100
			tax = taxRate
		}

		if earning > 30000 {
			balanceOfEarning = earning - 30000
			taxBalance = (30000 * 15) / 100
			balance = (balanceOfEarning * 20) / 100
			tax = taxBalance + balance
		}

		totalTax = tax + totalTax

		fmt.Printf("%.2f is the citizen's tax%n", tax)
	}

	fmt.Printf("%.2f is the citizen's totalTax%n", totalTax)

}
