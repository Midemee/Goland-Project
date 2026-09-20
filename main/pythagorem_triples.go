package main

//5.21
import (
	"fmt"
	"math"
)

func main() {
	const maxSide = 500

	fmt.Printf("%-8s%-8s%s\n", "side1", "side2", "hypotenuse")
	for side1 := 1; side1 <= maxSide; side1++ {
		for side2 := side1; side2 <= maxSide; side2++ {
			sumOfSquares := side1*side1 + side2*side2
			hypotenuse := int(math.Sqrt(float64(sumOfSquares)))

			if hypotenuse <= maxSide && hypotenuse*hypotenuse == sumOfSquares {
				fmt.Printf("%-8d%-8d%d\n", side1, side2, hypotenuse)
			}
		}
	}
}
