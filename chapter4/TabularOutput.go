package main

import "fmt"

func main() {

	fmt.Printf("%s %7s %6s %5s%n", "N", "N2", "N3", "N4")
	fmt.Printf("%d %6d %6d %5d%n", 1, 1, 1, 1)
	fmt.Printf("%d %6d %6d %6d%n", 2, 2*2, 2*2*2, 2*2*2*2)
	fmt.Printf("%d %6d %7d %5d%n", 3, 3*3, 3*3*3, 3*3*3*3)
	fmt.Printf("%d %7d %6d %6d%n", 4, 4*4, 4*4*4, 4*4*4*4)
	fmt.Printf("%d %7d %7d %5d%n", 5, 5*5, 5*5*5, 5*5*5*5)

}
