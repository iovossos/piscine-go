package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the full path of the program
	programPath := os.Args[0]

	// Initialize a variable to track the start of the program name
	start := 0

	// Find the position of the last '/' or '\' character
	for i, char := range programPath {
		if char == '/' || char == '\\' {
			start = i + 1
		}
	}

	// Extract the program name from the path
	programName := programPath[start:]

	// Print the program name using z01.PrintRune
	for _, char := range programName {
		z01.PrintRune(char)
	}
	// Print a newline
	z01.PrintRune('\n')
}
