package main

import "fmt"

// 5.16
func main() {
	var numbers [5]int

	for i := 0; i < 5; i++ {
		fmt.Printf("Enter number %d (1-30): ", i+1)
		fmt.Scan(&numbers[i])
	}

	fmt.Println("\nBar chart:")
	for i := 0; i < 5; i++ {
		for j := 1; j <= numbers[i]; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
