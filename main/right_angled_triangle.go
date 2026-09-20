package main

//4.30
import "fmt"

func main() {
	var base int

	fmt.Print("Enter the base length (1-10): ")
	fmt.Scan(&base)

	if base < 1 || base > 10 {
		fmt.Println("Invalid base length.")
		return
	}

	for row := 1; row <= base; row++ {

		for column := 1; column <= row; column++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
