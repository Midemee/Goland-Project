package main

//4.24
import "fmt"

func main() {
	var result int

	for i := 1; i <= 10; i++ {

		fmt.Print("Enter result (1 = pass, 2 = fail): ")
		fmt.Scan(&result)

		for result != 1 && result != 2 {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			fmt.Print("Enter result (1 = pass, 2 = fail): ")
			fmt.Scan(&result)
		}

		if result == 1 {
			fmt.Println("Pass")
		} else {
			fmt.Println("Fail")
		}
	}
}
