package main

//4.21
import "fmt"

func main() {
	var number int

	fmt.Print("Enter number 1: ")
	fmt.Scan(&number)

	largest := number

	for counter := 2; counter <= 10; counter++ {
		fmt.Printf("Enter number %d: ", counter)
		fmt.Scan(&number)

		if number > largest {
			largest = number
		}
	}

	fmt.Printf("The largest number is: %d\n", largest)
}
