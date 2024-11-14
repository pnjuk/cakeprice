# Cake Pricing Calculator

The **Cake Pricing Calculator** is a Go program that helps calculate the cost of cakes based on weight, base price per kilogram, and optional features like fillings and eatable prints. The user can specify a range of weights, as well as a custom weight, and get a detailed price breakdown for each option.

---

## Features

- **Weight Range Pricing**: Calculates prices for cake weights ranging from a minimum to a maximum value in 0.5 kg increments.
- **Custom Weight Pricing**: Allows users to input a custom weight and calculates its price separately.
- **Fillings with Custom Pricing**: Supports multiple cake fillings, each with its own price correction per kilogram.
- **Eatable Print Option**: Adds a fixed cost of 7 EUR for cakes with eatable prints.
- **Dynamic Input Validation**: Ensures all user inputs are valid and within acceptable ranges.

---

## How It Works

1. **User Input**: 
    - Minimum and maximum weights for the cake.
    - Base price per kilogram.
    - Whether the cake has various fillings.
    - Whether an eatable print is included.
    - Custom weight value (optional).

2. **Fillings and Corrections**:
    - If the user opts for fillings, default price corrections for different fillings are displayed.
    - Users can customize the corrections if needed.

3. **Price Calculation**:
    - Prices are calculated for each weight increment in the specified range.
    - If a custom weight is provided, its price is calculated and displayed separately after all other results.

4. **Output**:
    - Detailed prices are displayed, including the breakdown for each filling and weight combination.
    - Custom weight results are presented separately at the end.

---

## Installation
```bash
1. Clone the repository**:
   
git clone https://github.com/your-username/cake-pricing-calculator.git
Navigate to the project directory:

bash
Copy code
cd cake-pricing-calculator
Run the program:

bash
Copy code
go run main.go
```
## Usage
```
Enter the minimum and maximum weights:

Example: Minimum weight: 1, Maximum weight: 6
Provide the base price per kilogram:

Example: 15
Choose whether the cake has various fillings:

Input: yes or no
Select the eatable print option:

Input: yes or no
Enter custom weight (optional):

Example: 1.3
View the detailed price breakdown for each weight and filling.
```
## Example Output
```bash
Enter the minimal weight of the cake (e.g., 1): 1 /n
Enter the maximal weight of the cake (e.g., 6): 6
Enter the base price per kilo: 15
Do cake have various fillings? (yes/no): yes
Is there an eatable print option? (yes/no): yes
Do you have a custom weight value to calculate? (yes/no): yes
Enter your custom weight value: 1.3

Default corrections (additional price per kilo):
Filling: Raspberry, Correction: 1.00 EUR
Filling: Black currant, Correction: 0.00 EUR
Filling: Cherry, Correction: 0.00 EUR
Filling: Strawberry, Correction: 0.00 EUR

Are these corrections correct? (yes/no): yes

Base Price Calculation (for default weight steps):
Weight: 1.0 KG, Filling: Raspberry, Base Price: 15.00 EUR, Price with Print: 22.00 EUR
Weight: 1.0 KG, Filling: Black currant, Base Price: 15.00 EUR, Price with Print: 22.00 EUR
...

Base Price Calculation (for custom weight 1.3 KG):
Custom Weight: 1.3 KG, Filling: Raspberry, Base Price: 19.50 EUR, Price with Print: 26.50 EUR
Custom Weight: 1.3 KG, Filling: Black currant, Base Price: 19.50 EUR, Price with Print: 26.50 EUR
```
## Customization
Modify Fillings and Corrections:
The default filling corrections can be adjusted by the user during runtime.
Change Weight Steps:
The weight step is currently set to 0.5 kg. This can be modified in the code if needed.
