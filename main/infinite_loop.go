package main

//4.26
import "fmt"

func main() {
	row := 5
	column := 5

	for row >= 1 {

		for column >= 1 {
			if row%2 == 0 {
				fmt.Print("X")
			} else {
				fmt.Print("O")
			}

			column++
		}

		row--
		fmt.Println()
	}
}
