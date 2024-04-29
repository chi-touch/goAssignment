package main

import "fmt"

func main() {

	var number, counter, largest int64

	for count := 1; count <= 10; count++ {
		fmt.Print("Enter number: ")
		fmt.Scanln(number)

		if number > largest {
			largest = number
		}

		counter++
	}

	fmt.Printf("The winner of the sales contest is %d", largest)
}
