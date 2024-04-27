package main

import "fmt"

func main() {

	var count int
	var total, milesPerGallon, gallons, average, miles float64
	for miles != -1 && gallons != -1 {
		fmt.Print("Enter miles driven: ")
		fmt.Scanln(&miles)
		fmt.Print("Enter gallons: ")
		fmt.Scanln(&gallons)
		if miles != -1 && gallons != -1 {
			count++
			milesPerGallon = miles / gallons
			total += milesPerGallon
			average = float64(milesPerGallon) / float64(count)

			fmt.Printf("The miles/gallon is %f\n", milesPerGallon)
			fmt.Printf("%.2f is the average of milesPerGallon\n", average)
		}
	}
	fmt.Printf("%.2f is the total of the miles per gallon for the trip\n", total)
}
