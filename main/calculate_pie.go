package main

//5.20

import "fmt"

func main() {
	pi := 0.0
	sign := 1.0
	denominator := 1.0
	const totalTerms = 200000
	const target = "3.14159"

	termsNeeded := -1

	for term := 1; term <= totalTerms; term++ {
		pi += sign * 4.0 / denominator
		sign = -sign
		denominator += 2.0

		if termsNeeded == -1 && fmt.Sprintf("%.5f", pi) == target {
			termsNeeded = term
		}

		if term == 1 || term%20000 == 0 {
			fmt.Printf("Terms: %-10d Pi approximation: %.10f\n", term, pi)
		}
	}

	if termsNeeded != -1 {
		fmt.Printf("\nThe approximation first reaches %s after %d terms.\n", target, termsNeeded)
	} else {
		fmt.Printf("\nThe approximation never rounds to exactly %s within %d terms.\n", target, totalTerms)
	}
}
