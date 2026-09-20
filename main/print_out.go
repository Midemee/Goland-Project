package main

//5.19
import "fmt"

func main() {
	i, j, k, m := 2, 3, 2, 2

	fmt.Println("a)", i == 2)
	fmt.Println("b)", j == 5)
	fmt.Println("c)", (i >= 0) && (j <= 3))
	fmt.Println("d)", (m <= 100) && (k <= m))
	fmt.Println("e)", (j >= i) || (k != m))
	fmt.Println("f)", (k+i < j) || (4-j >= k))
	fmt.Println("g)", !(k > j))
}
