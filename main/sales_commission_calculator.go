package main

//4.19
import "fmt"

func main() {
	var itemValue float64
	totalSales := 0.0

	fmt.Println("Enter -1 when you are finished entering items")

	for {
		fmt.Print("Enter value of item sold: ")
		fmt.Scan(&itemValue)

		if itemValue == -1 {
			break
		}

		totalSales += itemValue
	}

	earnings := 200 + (totalSales * 0.09)

	fmt.Println("\nTotals Sales: $%.2f\n", totalSales)
	fmt.Println("Salesperson's earnings: $%.2f\n", earnings)
}
