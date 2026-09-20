package main

//4.34
import "fmt"

func main() {
	var target int
	var number int
	sum := 0

	fmt.Print("Enter the target sum: ")
	fmt.Scan(&target)

	for sum < target {
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)

		sum += number

		fmt.Println("Current sum:", sum)
	}

	fmt.Println("Target reached!")
}
