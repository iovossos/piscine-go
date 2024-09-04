package piscine

import (
	"github.com/01-edu/z01"
)

// PrintNbr prints the integer n as a string using z01.PrintRune.
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		printNumber(-n)
	} else {
		printNumber(n)
	}
}

// printNumber is a helper function that prints a non-negative integer.
func printNumber(n int) {
	if n >= 10 {
		printNumber(n / 10) // Recursively print the higher digits
	}
	z01.PrintRune(rune('0' + n%10)) // Print the current digit
}
