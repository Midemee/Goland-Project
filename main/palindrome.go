package main

//4.31
import "fmt"

func main() {
	var number int

	for {
		fmt.Print("Enter a five-digit integer: ")
		fmt.Scan(&number)

		if number >= 10000 && number <= 99999 {
			break
		}

		fmt.Println("Error: number must contain exactly five digits.")
	}

	original := number
	reversed := 0

	for number > 0 {
		digit := number % 10
		reversed = reversed*10 + digit
		number /= 10
	}

	if original == reversed {
		fmt.Println(original, "is a palindrome")
	} else {
		fmt.Println(original, "is not a palindrome")
	}
}
