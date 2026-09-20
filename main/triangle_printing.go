package main

//5.22
import "fmt"

const rows = 10

func main() {
	for i := 1; i <= rows; i++ {
		printStars(i)
		printSpaces(rows - i)
		fmt.Print("  ")

		printStars(rows - i + 1)
		printSpaces(i - 1)
		fmt.Print("  ")

		printSpaces(rows - i)
		printStars(i)
		fmt.Print("  ")

		printSpaces(i - 1)
		printStars(rows - i + 1)

		fmt.Println()
	}
}

func printStars(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
}

func printSpaces(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}
