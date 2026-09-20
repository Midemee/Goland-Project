package main

//5.24
import "fmt"

func main() {
	const size = 7 // width of the diamond at its widest row (must be odd)
	mid := size/2 + 1

	for i := 1; i <= size; i++ {
		var stars int
		if i <= mid {
			stars = 2*i - 1
		} else {
			stars = 2*(size-i+1) - 1
		}
		spaces := (size - stars) / 2

		for s := 1; s <= spaces; s++ {
			fmt.Print(" ")
		}
		for s := 1; s <= stars; s++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
