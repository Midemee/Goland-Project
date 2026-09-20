package main

//4.37
import "fmt"

func main() {
	var x1, y1, x2, y2 float64

	fmt.Print("Enter x1 and y1: ")
	fmt.Scan(&x1, &y1)

	fmt.Print("Enter x2 and y2: ")
	fmt.Scan(&x2, &y2)

	if x1 == x2 {
		fmt.Println("The points are on a vertical line.")
		fmt.Println("The line is perpendicular to the x-axis.")
	} else if y1 == y2 {
		fmt.Println("The points are on a horizontal line.")
		fmt.Println("The line is perpendicular to the y-axis.")
	} else {
		fmt.Println("The line is not perpendicular to an axis.")
	}
}
