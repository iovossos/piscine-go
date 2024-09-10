package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	// Check for base validity
	if !isValidBase(base) {
		printString("NV")
		return
	}

	baseLen := len(base)

	// Handle negative numbers
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	// Convert the number to the base and print it
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	var result []rune
	for nbr > 0 {
		result = append([]rune{rune(base[nbr%baseLen])}, result...)
		nbr /= baseLen
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
}

// Helper function to check base validity
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Helper function to print a string using z01.PrintRune
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
