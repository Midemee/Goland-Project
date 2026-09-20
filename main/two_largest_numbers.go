package main

//4.23
import "fmt"

func main() {
	var number int
	var largest, secondLargest int

	fmt.Print("Enter number 1: ")
	fmt.Scan(&number)

	largest = number
	secondLargest = number

	for i := 2; i <= 10; i++ {
		fmt.Printf("Enter number %d: ", i)
		fmt.Scan(&number)

		if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest {
			secondLargest = number
		}
	}

	fmt.Println("Largest:", largest)
	fmt.Println("Second largest:", secondLargest)
}
