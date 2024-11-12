package main

import (
	"fmt"
)

func main() {
	var hasFillings string
	var minWeight float64

	// Ask for the minimal weight of the cake
	fmt.Print("Enter the minimal weight of the cake (e.g., 1.5): ")
	_, err := fmt.Scanf("%f", &minWeight)
	if err != nil || minWeight < 1.5 {
		fmt.Println("Invalid input. Please enter a valid weight (minimum 1.5).")
		return
	}

	// Ask if the cake has various fillings
	fmt.Print("Do cake have various fillings? (yes/no): ")
	fmt.Scanf("%s", &hasFillings)

	if hasFillings == "no" {
		var price float64

		// Ask for price input for cakes without fillings
		fmt.Print("Enter the price for the cake (without fillings): ")
		_, err := fmt.Scanf("%f", &price)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		// Use initial logic for cakes without fillings
		for multiplier := minWeight; multiplier <= 6; multiplier += 0.5 {
			result := price * multiplier
			printYes := result + 7

			// Print the weight and results
			fmt.Printf("Weight: %.1f KG, Price without print: %.2f, Print yes: %.2f\n", multiplier, result, printYes)
		}
	} else if hasFillings == "yes" {
		// Define the default prices for fillings
		fillings := map[string]float64{
			"Raspberry":    42.0,
			"Black currant": 39.0,
			"Cherry":        39.0,
			"Strawberry":    39.0,
		}

		// Display default prices for 1 kg and 1.5 kg
		fmt.Println("\nDefault prices:")
		for filling, fillingPrice := range fillings {
			fmt.Printf("Filling: %s, Price for 1 KG: %.2f EUR, Price for 1.5 KG: %.2f EUR\n", filling, fillingPrice/1.5, fillingPrice)
		}

		// Ask user if the default prices are correct
		var priceCorrect string
		fmt.Print("\nAre these prices correct? (yes/no): ")
		fmt.Scanf("%s", &priceCorrect)

		if priceCorrect == "no" {
			// Allow user to enter custom prices for 1.5 KG for each filling
			for filling := range fillings {
				var customPrice float64
				fmt.Printf("Enter custom price for 1.5 KG of %s: ", filling)
				_, err := fmt.Scanf("%f", &customPrice)
				if err != nil {
					fmt.Println("Invalid input. Please enter a valid number.")
					return
				}
				fillings[filling] = customPrice
			}
		}

		// Calculate prices for each filling from minWeight to 4 KG
		for weight := minWeight; weight <= 4; weight += 0.5 {
			for filling, fillingPrice := range fillings {
				pricePerKilo := fillingPrice / 1.5
				totalPrice := pricePerKilo * weight
				fmt.Printf("Weight: %.1f KG, Filling: %s, Total Price: %.2f EUR\n", weight, filling, totalPrice)
			}
		}
	} else {
		fmt.Println("Invalid input for fillings. Please enter 'yes' or 'no'.")
	}
}

