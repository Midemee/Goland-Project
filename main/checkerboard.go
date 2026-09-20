package main

//4.33
import "fmt"

func main() {
	for row := 1; row <= 8; row++ {

		if row%2 == 0 {
			fmt.Print(" ")
		}

		for column := 1; column <= 8; column++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}
