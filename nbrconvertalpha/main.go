package main

import (
	"os"
	"piscine" // Assuming this is where your atoi function is located

	"github.com/01-edu/z01"
)

func main() {
	// Check if the --upper flag is provided
	upper := false
	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		upper = true
		os.Args = os.Args[1:] // Skip the flag for further processing
	}

	// Loop through each argument (excluding the program name)
	for i := 1; i < len(os.Args); i++ {
		// Convert the argument to an integer using the atoi function from "piscine"
		n := piscine.Atoi(os.Args[i])
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
