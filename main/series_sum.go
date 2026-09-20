package main

//5.13
import "fmt"

func main() {
	var sum int64 = 0

	fmt.Printf("%-5s%s\n", "n", "sum")
	for n := int64(1); n <= 100; n++ {
		sum += n
		fmt.Printf("%-5d%d\n", n, sum)
	}
}
