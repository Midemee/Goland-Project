package main

//5.23
import "fmt"

func main() {
	fmt.Println("Checking: !(x<5) && !(y>=7)  ==  !((x<5) || (y>=7))")
	mismatchA := false
	for x := -3; x <= 10; x++ {
		for y := -3; y <= 10; y++ {
			left := !(x < 5) && !(y >= 7)
			right := !((x < 5) || (y >= 7))
			if left != right {
				mismatchA = true
			}
		}
	}
	fmt.Println("All cases match:", !mismatchA)

	fmt.Println("\nChecking: !(a==b) || !(g!=5)  ==  !((a==b) && (g!=5))")
	mismatchB := false
	for a := 0; a <= 5; a++ {
		for b := 0; b <= 5; b++ {
			for g := 0; g <= 8; g++ {
				left := !(a == b) || !(g != 5)
				right := !((a == b) && (g != 5))
				if left != right {
					mismatchB = true
				}
			}
		}
	}
	fmt.Println("All cases match:", !mismatchB)
}
