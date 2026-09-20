package main

//5.14
import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.00

	for ratePercent := 5; ratePercent <= 10; ratePercent++ {
		rate := float64(ratePercent) / 100.0

		fmt.Printf("Interest Rate: %d%%\n", ratePercent)
		fmt.Printf("%-6s%s\n", "Year", "Amount on deposit")

		for year := 1; year <= 10; year++ {
			amount := principal * math.Pow(1.0+rate, float64(year))
			fmt.Printf("%-6d%.2f\n", year, amount)
		}
		fmt.Println()
	}
}
