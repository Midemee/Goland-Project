package main

//5.15
import "fmt"

func main() {
	const rows = 10

	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 1; i <= rows; i++ {
		for s := 1; s <= rows-i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := rows; i >= 1; i-- {
		for s := 1; s <= rows-i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
