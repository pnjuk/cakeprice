package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var hasFillings, hasEatablePrint, hasCustomWeights string
	var minWeight, maxWeight, basePrice float64
	var customWeights []float64

	reader := bufio.NewReader(os.Stdin)

	// Function to get float input with validation
	getFloatInput := func(prompt string, min, max float64) (float64, error) {
		var input float64
		fmt.Print(prompt)
		_, err := fmt.Scanf("%f", &input)
		if err != nil || input < min || (max > 0 && input > max) {
			return 0, fmt.Errorf("invalid input")
		}
		return input, nil
	}

	// Function to get string input
	getStringInput := func(prompt string) string {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}

	// Get minimal weight
	minWeight, err := getFloatInput("Enter the minimal weight of the cake (e.g., 1.0): ", 1.0, 0)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid weight (minimum 1.0).")
		return
	}

	// Get maximal weight
	maxWeight, err = getFloatInput("Enter the maximal weight of the cake (e.g., 6): ", minWeight, 6)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid weight (maximum up to 6 and greater than or equal to the minimum weight).")
		return
	}

	// Get base price per kilo
	basePrice, err = getFloatInput("Enter the base price per kilo: ", 0.01, 0)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid price.")
		return
	}

	// Get other inputs
	hasFillings = getStringInput("Do cake have various fillings? (yes/no): ")
	hasEatablePrint = getStringInput("Is there an eatable print option? (yes/no): ")
	addPrintCost := hasEatablePrint == "yes"
	hasCustomWeights = getStringInput("Do you have custom weights? (yes/no): ")

	if hasCustomWeights == "yes" {
		for {
			fmt.Print("Enter a custom weight (or press Enter to finish): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "" {
				break
			}
			customWeight, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid custom weight.")
				continue
			}
			customWeights = append(customWeights, customWeight)
		}
	}

	fillings := map[string]float64{}
	if hasFillings == "yes" {
		fillings = map[string]float64{
			"Raspberry":     1.0,
			"Black currant": 0.0,
			"Cherry":        0.0,
			"Strawberry":    0.0,
		}

		fmt.Println("\nDefault corrections (additional price per kilo):")
		for filling, correction := range fillings {
			fmt.Printf("Filling: %s, Correction: %.2f EUR\n", filling, correction)
		}

		if getStringInput("\nAre these corrections correct? (yes/no): ") == "no" {
			for filling := range fillings {
				correction, err := getFloatInput(fmt.Sprintf("Enter custom correction for %s: ", filling), 0, 0)
				if err != nil {
					fmt.Println("Invalid input. Please enter a valid number.")
					return
				}
				fillings[filling] = correction
			}
		}
	}

	// Function to calculate and print prices
	calculatePrices := func(weight float64) {
		baseWeightPrice := basePrice * weight
		if hasFillings == "yes" {
			for filling, correction := range fillings {
				finalPrice := baseWeightPrice + (correction * weight)
				if addPrintCost {
					fmt.Printf("Weight: %.1f KG, Filling: %s, Base Price: %.2f EUR, Price with Print: %.2f EUR\n", weight, filling, finalPrice, finalPrice+7)
				} else {
					fmt.Printf("Weight: %.1f KG, Filling: %s, Base Price: %.2f EUR\n", weight, filling, finalPrice)
				}
			}
		} else {
			if addPrintCost {
				fmt.Printf("Weight: %.1f KG, Base Price: %.2f EUR, Price with Print: %.2f EUR\n", weight, baseWeightPrice, baseWeightPrice+7)
			} else {
				fmt.Printf("Weight: %.1f KG, Base Price: %.2f EUR\n", weight, baseWeightPrice)
			}
		}
	}

	// Main block of prices
	fmt.Println("\nMain block of prices:")

	// Calculate prices for the range of weights
	for weight := minWeight; weight <= maxWeight; weight += 0.5 {
		calculatePrices(weight)
	}

	// Calculate prices for the custom weights if provided
	if hasCustomWeights == "yes" {
		fmt.Println("\nCustom Weight Prices:")
		for _, customWeight := range customWeights {
			calculatePrices(customWeight)
		}
	}
}
