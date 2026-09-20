package main

//4.22
import "fmt"

func main() {
	fmt.Printf("%-5s %-5s %-5s %-5s\n", "N", "N²", "N³", "N⁴")

	for n := 1; n <= 5; n++ {
		n2 := n * n
		n3 := n * n * n
		n4 := n * n * n * n

		fmt.Printf("%-5d %-5d %-5d %-5d\n", n, n2, n3, n4)
	}
}
