package main

//4.20
import "fmt"

func main() {
	const taxCeiling = 30000.0
	const lowerTaxRate = 0.15
	const higherTaxRate = 0.20

	var name string
	var earnings float64

	for citizen := 1; citizen <= 3; citizen++ {
		fmt.Printf("\nEnter name of citizen %d: ", citizen)
		fmt.Scan(&name)

		fmt.Print("Enter yearly earnings: ")
		fmt.Scan(&earnings)

		var tax float64

		if earnings <= taxCeiling {
			tax = earnings * lowerTaxRate
		} else {
			tax = (taxCeiling * lowerTaxRate) +
				((earnings - taxCeiling) * higherTaxRate)
		}

		fmt.Printf("Citizen: %s\n", name)
		fmt.Printf("Total tax: $%.2f\n", tax)
	}
}
