package main

//5.11
import "fmt"

func main() {
	var count int
	fmt.Print("How many values will you enter? ")
	fmt.Scan(&count)

	var min, max, num int
	fmt.Printf("Enter %d integers, one at a time:\n", count)

	fmt.Scan(&num)
	min, max = num, num

	for i := 1; i < count; i++ {
		fmt.Scan(&num)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Printf("\nMinimum: %d\n", min)
	fmt.Printf("Maximum: %d\n", max)
	fmt.Printf("Sum of the two extremes: %d\n", min+max)
}
