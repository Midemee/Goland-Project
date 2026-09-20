package main

//5.12
import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			sum += i
		}
	}
	fmt.Printf("Sum of integers between 1 and 30 that are divisible by 3: %d\n", sum)
}
