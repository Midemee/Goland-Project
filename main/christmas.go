package main

import "fmt"

func main() {
	for day := 1; day <= 12; day++ {
		fmt.Print("On the ")
		printDayName(day)
		fmt.Println(" day of Christmas")
		fmt.Println("My true love sent to me:")

		for line := day; line >= 1; line-- {
			printGiftLine(line, day)
		}
		fmt.Println()
	}
}

func printDayName(day int) {
	switch day {
	case 1:
		fmt.Print("first")
	case 2:
		fmt.Print("second")
	case 3:
		fmt.Print("third")
	case 4:
		fmt.Print("fourth")
	case 5:
		fmt.Print("fifth")
	case 6:
		fmt.Print("sixth")
	case 7:
		fmt.Print("seventh")
	case 8:
		fmt.Print("eighth")
	case 9:
		fmt.Print("ninth")
	case 10:
		fmt.Print("tenth")
	case 11:
		fmt.Print("eleventh")
	case 12:
		fmt.Print("twelfth")
	}
}

func printGiftLine(line int, day int) {
	switch line {
	case 1:
		if day == 1 {
			fmt.Println("A partridge in a pear tree.")
		} else {
			fmt.Println("And a partridge in a pear tree.")
		}
	case 2:
		fmt.Println("Two turtle doves,")
	case 3:
		fmt.Println("Three French hens,")
	case 4:
		fmt.Println("Four calling birds,")
	case 5:
		fmt.Println("Five golden rings,")
	case 6:
		fmt.Println("Six geese a-laying,")
	case 7:
		fmt.Println("Seven swans a-swimming,")
	case 8:
		fmt.Println("Eight maids a-milking,")
	case 9:
		fmt.Println("Nine ladies dancing,")
	case 10:
		fmt.Println("Ten lords a-leaping,")
	case 11:
		fmt.Println("Eleven pipers piping,")
	case 12:
		fmt.Println("Twelve drummers drumming,")
	}
}
