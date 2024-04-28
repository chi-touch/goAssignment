package main

import (
	"fmt"
	"math"
)

func main() {

	var commission, itemValue, value, salesPersonEarnings float64

	for count := 0; count < math.MaxInt; count++ {

		fmt.Println("Enter the value of the item: ")
		fmt.Scanln(&itemValue)
		value += itemValue

		commission = 0.9 * value

		fmt.Println("Do you want to add another item")
		var prompt string
		fmt.Scanln(&prompt)
		if prompt == "yes" {
			continue
		} else {
			break
		}
	}
	salesPersonEarnings = 200 + commission

	fmt.Println("this is your weekly wage", salesPersonEarnings)

}
