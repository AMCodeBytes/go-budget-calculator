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
			fmt.Println("-----   Potential Earnings   -----")
			potentialEarnings()
		case 2:
			fmt.Println("-----   Potential Earnings Per Year   -----")
			perYear()
		case 3:
			fmt.Println("-----   Potential Earnings Per Year   -----")
			perMonth()
		case 4:
			fmt.Println("-----   Potential Earnings Per Year   -----")
			perWeek()
		case 5:
			// code
		case 6:
			os.Exit(0)
		default:
			panic("Invalid input entered.")
		}

		choice = 0
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

func potentialEarnings() {
	var hours int
	var pay float64
	var result float64

	fmt.Print("How many hours: ")
	fmt.Scan(&hours)

	fmt.Print("How much pay: ")
	fmt.Scan(&pay)

	result = float64(hours) * pay

	fmt.Printf("\n\n----- Result -----\nFor %v hours at £%v pay you'll make £%v\n\n", hours, pay, result)
}

func perYear() {
	var pay float64
	var result float64

	fmt.Print("How much pay: ")
	fmt.Scan(&pay)

	result = pay * float64(8) * (float64(260))

	fmt.Printf("\n\n----- Result -----\nFor a year at £%v pay you'll make £%v\n\n", pay, result)
}

func perMonth() {
	var pay float64
	var result float64

	fmt.Print("How much pay: ")
	fmt.Scan(&pay)

	result = pay * float64(8) * (float64(21))

	fmt.Printf("\n\n----- Result -----\nFor a month at £%v pay you'll make £%v\n\n", pay, result)
}

func perWeek() {
	var pay float64
	var result float64

	fmt.Print("How much pay: ")
	fmt.Scan(&pay)

	result = pay * float64(8) * (float64(5))

	fmt.Printf("\n\n----- Result -----\nFor a week at £%v pay you'll make £%v\n\n", pay, result)
}
