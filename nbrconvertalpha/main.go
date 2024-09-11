package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Custom atoi function to convert string to integer
func atoi(s string) int {
	result := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			// Return 0 for invalid characters (non-digit)
			return 0
		}
		result = result*10 + int(char-'0')
	}
	return result
}

func main() {
	// Check if the --upper flag is provided
	upper := false
	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		upper = true
		os.Args = os.Args[1:] // Skip the flag for further processing
	}

	// If no arguments are provided, just print a newline and exit
	if len(os.Args) == 1 {
		z01.PrintRune('\n')
		return
	}

	// Loop through each argument (excluding the program name)
	for i := 1; i < len(os.Args); i++ {
		// Convert the argument to an integer using the custom atoi function
		n := atoi(os.Args[i])
		if n < 1 || n > 26 {
			// Print a space if the argument is not valid or not in the range 1-26
			z01.PrintRune(' ')
			continue
		}

		// Calculate the corresponding letter
		var letter rune
		if upper {
			letter = rune('A' + n - 1) // Uppercase letter
		} else {
			letter = rune('a' + n - 1) // Lowercase letter
		}

		// Print the letter
		z01.PrintRune(letter)
	}

	// Print a newline at the end
	z01.PrintRune('\n')
}
