package main

//4.18
import "fmt"

func main() {
	var miles int
	var gallons int

	totalMiles := 0
	totalGallons := 0

	fmt.Println("Enter -1 for miles to end the program")

	for {
		fmt.Print("Enter miles driven: ")
		fmt.Scan(&miles)

		if miles == -1 {
			break
		}

		fmt.Print("Enter gallons used: ")
		fmt.Scan(&gallons)

		milesPerGallon := float64(miles) / float64(gallons)

		totalMiles += miles
		totalGallons += gallons

		combinedMPG := float64(totalMiles) / float64(totalGallons)

		fmt.Printf("Miles per gallon for this trip: %.2f\n", milesPerGallon)
		fmt.Printf("Combined miles per gallon: %.2f\n\n", combinedMPG)
	}
}
