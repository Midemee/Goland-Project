package main

//4.17
import "fmt"

func main() {
	var accountNumber int
	var beginningBalance int
	var charges int
	var credits int
	var creditLimit int

	fmt.Println("Enter -1 for account number to end the program.")

	for {
		fmt.Print("Enter account number: ")
		fmt.Scan(&accountNumber)

		if accountNumber == -1 {
			break
		}

		fmt.Print("Enter beginning balance: ")
		fmt.Scan(&beginningBalance)

		fmt.Print("Enter total charges: ")
		fmt.Scan(&charges)

		fmt.Print("Enter total credits: ")
		fmt.Scan(&credits)

		fmt.Print("Enter credit limit: ")
		fmt.Scan(&creditLimit)

		newBalance := beginningBalance + charges - credits

		fmt.Printf("\nAccount number: %d\n", accountNumber)
		fmt.Printf("New balance: %d\n", newBalance)
		fmt.Printf("Credit limit: %d\n", creditLimit)

		if newBalance > creditLimit {
			fmt.Println("Credit limit exceeded")
		}

		fmt.Println()
	}
}
