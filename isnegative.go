package piscine

import (
	"github.com/01-edu/z01"
)

// IsNegative prints 'T' if the integer nb is negative,
// otherwise it prints 'F'.
func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n') // Print a newline for better output formatting
}
