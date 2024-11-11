package main

import (
	"fmt"
)

func main() {
	var price float64

	// Ask for price input from the user
	fmt.Print("Enter the price: ")
	_, err := fmt.Scanf("%f", &price)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	// Loop through multipliers from 1.5 to 6 with a step of 0.5
	for multiplier := 1.5; multiplier <= 6; multiplier += 0.5 {
		result := price * multiplier
		printYes := result + 7

		// Print the results
		fmt.Printf("Result: %.2f, Print yes: %.2f\n", result, printYes)
	}
}
