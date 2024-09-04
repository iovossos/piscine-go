package piscine

import (
	"github.com/01-edu/z01"
)

// PrintComb2 prints all possible combinations of two different two-digit numbers
// in ascending order, with each combination separated by a comma and a space.
func PrintComb2() {
	for i := 10; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			printTwoDigitNumber(i)
			printTwoDigitNumber(j)
			if !(i == 98 && j == 99) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n') // Print a newline at the end
}

// printTwoDigitNumber prints a two-digit number using z01.PrintRune
func printTwoDigitNumber(n int) {
	// Print the tens place
	z01.PrintRune(rune(n/10 + '0'))
	// Print the units place
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	PrintComb2()
}
