package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Loop through all the arguments except the program name (os.Args[0])
	for i := 1; i < len(os.Args); i++ {
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
