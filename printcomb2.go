package piscine

import (
	"github.com/01-edu/z01"
)

// PrintComb2 prints all unique combinations of two different two-digit numbers.
func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			printTwoDigit(i)
			z01.PrintRune(' ')
			printTwoDigit(j)

			if !(i == 98 && j == 99) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n') // Print newline at the end
}

// Helper function to print a two-digit number
func printTwoDigit(n int) {
	// Print tens digit
	z01.PrintRune(rune('0' + n/10))
	// Print units digit
	z01.PrintRune(rune('0' + n%10))
}
