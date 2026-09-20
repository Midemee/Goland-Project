package main

//5.26
import "fmt"

func main() {
	var n int
	for {
		fmt.Print("Enter an odd number between 1 and 19: ")
		fmt.Scan(&n)
		if n >= 1 && n <= 19 && n%2 != 0 {
			break
		}
		fmt.Println("Invalid input. Please enter an odd number between 1 and 19.")
	}

	mid := n/2 + 1
	for i := 1; i <= n; i++ {
		var stars int
		if i <= mid {
			stars = 2*i - 1
		} else {
			stars = 2*(n-i+1) - 1
		}
		spaces := (n - stars) / 2

		for s := 1; s <= spaces; s++ {
			fmt.Print(" ")
		}
		for s := 1; s <= stars; s++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
