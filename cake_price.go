package main

import (
	"fmt"
)

func main() {
	var hasFillings string
	var minWeight, maxWeight float64
	var hasEatablePrint string

	// Ask for the minimal weight of the cake
	fmt.Print("Enter the minimal weight of the cake (e.g., 1.5): ")
	_, err := fmt.Scanf("%f", &minWeight)
	if err != nil || minWeight < 1.5 {
		fmt.Println("Invalid input. Please enter a valid weight (minimum 1.5).")
		return
	}

	// Ask for the maximal weight of the cake
	fmt.Print("Enter the maximal weight of the cake (e.g., 6): ")
	_, err = fmt.Scanf("%f", &maxWeight)
	if err != nil || maxWeight > 6 || maxWeight < minWeight {
		fmt.Println("Invalid input. Please enter a valid weight (maximum up to 6 and greater than or equal to the minimum weight).")
		return
	}

	// Ask if the cake has various fillings
	fmt.Print("Do cake have various fillings? (yes/no): ")
	fmt.Scanf("%s", &hasFillings)

	// Ask if there's an eatable print option
	fmt.Print("Is there an eatable print option? (yes/no): ")
	fmt.Scanf("%s", &hasEatablePrint)
	addPrintCost := hasEatablePrint == "yes"

	if hasFillings == "no" {
		var price float64

		// Ask for price input for cakes without fillings
		fmt.Print("Enter the price for the cake (without fillings): ")
		_, err := fmt.Scanf("%f", &price)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		// Calculate prices for cakes without fillings
		for multiplier := minWeight; multiplier <= maxWeight; multiplier += 0.5 {
			basePrice := price * multiplier
			if addPrintCost {
				priceWithPrint := basePrice + 7
				fmt.Printf("Weight: %.1f KG, Base Price: %.2f, Price with Print: %.2f\n", multiplier, basePrice, priceWithPrint)
			} else {
				fmt.Printf("Weight: %.1f KG, Base Price: %.2f\n", multiplier, basePrice)
			}
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

		// Calculate prices for each filling from minWeight to maxWeight
		for weight := minWeight; weight <= maxWeight; weight += 0.5 {
			for filling, fillingPrice := range fillings {
				pricePerKilo := fillingPrice / 1.5
				basePrice := pricePerKilo * weight
				if addPrintCost {
					priceWithPrint := basePrice + 7
					fmt.Printf("Weight: %.1f KG, Filling: %s, Base Price: %.2f EUR, Price with Print: %.2f EUR\n", weight, filling, basePrice, priceWithPrint)
				} else {
					fmt.Printf("Weight: %.1f KG, Filling: %s, Base Price: %.2f EUR\n", weight, filling, basePrice)
				}
			}
		}
	} else {
		fmt.Println("Invalid input for fillings. Please enter 'yes' or 'no'.")
	}
}
