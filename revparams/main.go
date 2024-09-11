package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Loop through the arguments in reverse, starting from the last one
	for i := len(os.Args) - 1; i > 0; i-- {
		// Get the current argument
		arg := os.Args[i]

		// Print each character of the argument
		for _, char := range arg {
			z01.PrintRune(char)
		}

		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
