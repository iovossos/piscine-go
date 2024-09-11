package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the command-line arguments (excluding the program name)
	args := os.Args[1:]

	// Perform a simple bubble sort to sort the arguments in ASCII order
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				// Swap args[i] and args[j]
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	// Print each argument in the sorted order
	for _, arg := range args {
		// Print each character of the argument
		for _, char := range arg {
			z01.PrintRune(char)
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
