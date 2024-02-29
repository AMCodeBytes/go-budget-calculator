package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int

	fmt.Println("-----   Budget Calculator   -----")

	for {
		displayMenu()

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// code
		case 2:
			// code
		case 3:
			// code
		case 4:
			// code
		case 5:
			// code
		case 6:
			os.Exit(0)
		default:
			panic("Invalid input entered.")
		}
	}
}

func displayMenu() {
	fmt.Println("1. Calculate to find out how much you will make from x hours with y pay")
	fmt.Println("2. Calculate how much a year (approx) you can make from y pay")
	fmt.Println("3. Calculate how much a month (approx) you can make from y pay")
	fmt.Println("4. Calculate how much a week (approx) you can make from y pay")
	fmt.Println("5. Calculate how much a day (approx) you can make from y pay")
	fmt.Println("6. Quit")
}
