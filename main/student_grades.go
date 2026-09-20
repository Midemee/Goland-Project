package main

//5.17
import "fmt"

func main() {
	var countA, countB, countC, countD, countF int

	for student := 1; student <= 5; student++ {
		var name, grade string
		fmt.Printf("Enter name for student %d: ", student)
		fmt.Scan(&name)
		fmt.Printf("Enter letter grade for %s: ", name)
		fmt.Scan(&grade)

		switch grade {
		case "A", "a":
			countA++
		case "B", "b":
			countB++
		case "C", "c":
			countC++
		case "D", "d":
			countD++
		case "F", "f":
			countF++
		default:
			fmt.Println("Invalid grade entered.")
		}
	}

	fmt.Println("\nNumber of students who received each grade:")
	fmt.Printf("A: %d\n", countA)
	fmt.Printf("B: %d\n", countB)
	fmt.Printf("C: %d\n", countC)
	fmt.Printf("D: %d\n", countD)
	fmt.Printf("F: %d\n", countF)
}
