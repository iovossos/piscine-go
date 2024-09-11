package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printLetter(n int, upper bool) {
	if n < 1 || n > 26 {
		z01.PrintRune(' ')
		return
	}
	letter := rune(n + 96) // 'a' is at position 97 in Unicode
	if upper {
		letter -= 32 // Convert to uppercase
	}
	z01.PrintRune(letter)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		return
	}

	upper := false
	if args[1] == "--upper" {
		upper = true
		args = args[1:]
	}

	// Flag to check if any valid output was produced
	hasOutput := false

	for _, arg := range args[1:] {
		n := 0
		valid := true
		for _, ch := range arg {
			if ch < '0' || ch > '9' {
				z01.PrintRune(' ')
				valid = false
				break
			}
			n = n*10 + int(ch-'0')
		}
		if valid {
			printLetter(n, upper)
			hasOutput = true
		}
	}

	// Print a newline if any valid output was produced
	if hasOutput {
		z01.PrintRune('\n')
	}
}
