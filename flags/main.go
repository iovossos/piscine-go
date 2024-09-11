package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Function to print help message
func printHelp() {
	helpText := `--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`

	for _, ch := range helpText {
		z01.PrintRune(ch)
	}
}

// Function to check if a string starts with a prefix
func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

// Function to split a string by '='
func splitByEquals(s string) (string, string) {
	for i, ch := range s {
		if ch == '=' {
			return s[:i], s[i+1:]
		}
	}
	return s, ""
}

// Function to print a string
func printString(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	insertStr := ""
	order := false
	var inputStr string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if hasPrefix(arg, "--insert=") {
			_, insertStr = splitByEquals(arg)
		} else if hasPrefix(arg, "-i=") {
			_, insertStr = splitByEquals(arg)
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			inputStr = arg
		}
	}

	if inputStr == "" {
		printHelp()
		return
	}

	if order {
		// Convert string to rune slice for sorting
		runes := []rune(inputStr)
		for i := 0; i < len(runes); i++ {
			for j := i + 1; j < len(runes); j++ {
				if runes[i] > runes[j] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		inputStr = string(runes)
	}

	if insertStr != "" {
		inputStr += insertStr
	}

	printString(inputStr)
	z01.PrintRune('\n') // Print a newline at the end
}
