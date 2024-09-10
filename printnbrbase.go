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

	// Special case for the minimum int value (math.MinInt64)
	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		printString(convertMinInt64(base, baseLen))
		return
	}

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

// Handle the minimum value for int64
func convertMinInt64(base string, baseLen int) string {
	minInt := -9223372036854775808
	var result []rune

	for minInt < 0 {
		result = append([]rune{rune(base[-(minInt % baseLen)])}, result...)
		minInt /= baseLen
	}

	return string(result)
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
