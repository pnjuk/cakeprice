package main

import (
	"fmt"
)

func main() {
	var hasFillings string
	var minWeight, maxWeight, basePrice float64
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

	// Ask for the base price per kilo
	fmt.Print("Enter the base price per kilo: ")
	_, err = fmt.Scanf("%f", &basePrice)
	if err != nil || basePrice <= 0 {
		fmt.Println("Invalid input. Please enter a valid price.")
		return
	}

	// Ask if the cake has various fillings
	fmt.Print("Do cake have various fillings? (yes/no): ")
	fmt.Scanf("%s", &hasFillings)

	// Ask if there's an eatable print option
	fmt.Print("Is there an eatable print option? (yes/no): ")
	fmt.Scanf("%s", &hasEatablePrint)
	addPrintCost := hasEatablePrint == "yes"

	fillings := map[string]float64{}
	if hasFillings == "yes" {
		// Set default corrections
		fillings = map[string]float64{
			"Raspberry":     1.0,
			"Black currant": 0.0,
			"Cherry":        0.0,
			"Strawberry":    0.0,
		}

		// Display default corrections
		fmt.Println("\nDefault corrections (additional price per kilo):")
		for filling, correction := range fillings {
			fmt.Printf("Filling: %s, Correction: %.2f EUR\n", filling, correction)
		}

		// Ask user if the corrections are correct
		var correctionsCorrect string
		fmt.Print("\nAre these corrections correct? (yes/no): ")
		fmt.Scanf("%s", &correctionsCorrect)

		if correctionsCorrect == "no" {
			// Allow user to enter custom corrections for each filling
			for filling := range fillings {
				var customCorrection float64
				fmt.Printf("Enter custom correction for %s: ", filling)
				_, err := fmt.Scanf("%f", &customCorrection)
				if err != nil {
					fmt.Println("Invalid input. Please enter a valid number.")
					return
				}
				fillings[filling] = customCorrection
			}
		}
	}

	// Calculate prices
	for weight := minWeight; weight <= maxWeight; weight += 0.5 {
		baseWeightPrice := basePrice * weight

		if hasFillings == "yes" {
			for filling, correction := range fillings {
				finalPrice := baseWeightPrice + (correction * weight)
				if addPrintCost {
					priceWithPrint := finalPrice + 7
					fmt.Printf("Weight: %.1f KG, Filling: %s, Base Price: %.2f EUR, Price with Print: %.2f EUR\n", weight, filling, finalPrice, priceWithPrint)
				} else {
					fmt.Printf("Weight: %.1f KG, Filling: %s, Base Price: %.2f EUR\n", weight, filling, finalPrice)
				}
			}
		} else {
			if addPrintCost {
				priceWithPrint := baseWeightPrice + 7
				fmt.Printf("Weight: %.1f KG, Base Price: %.2f EUR, Price with Print: %.2f EUR\n", weight, baseWeightPrice, priceWithPrint)
			} else {
				fmt.Printf("Weight: %.1f KG, Base Price: %.2f EUR\n", weight, baseWeightPrice)
			}
		}
	}
}
