package main

//4.32
import "fmt"

func main() {
	var binary int

	fmt.Print("Enter a binary number: ")
	fmt.Scan(&binary)

	original := binary
	decimal := 0
	placeValue := 1

	for binary > 0 {
		digit := binary % 10

		if digit != 0 && digit != 1 {
			fmt.Println("Invalid binary number.")
			return
		}

		decimal += digit * placeValue

		placeValue *= 2
		binary /= 10
	}

	fmt.Println("Decimal equivalent of", original, "is", decimal)
}
