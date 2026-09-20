package main

//5.18
import "fmt"

func main() {
	for ratePercent := 5; ratePercent <= 10; ratePercent++ {
		var amountPennies int64 = 100000

		fmt.Printf("Interest Rate: %d%%\n", ratePercent)
		fmt.Printf("%-6s%s\n", "Year", "Amount on deposit")

		for year := 1; year <= 10; year++ {
			amountPennies = (amountPennies*(100+int64(ratePercent)) + 50) / 100

			dollars := amountPennies / 100
			cents := amountPennies % 100
			fmt.Printf("%-6d$%d.%02d\n", year, dollars, cents)
		}
		fmt.Println()
	}
}
