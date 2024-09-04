package main

import ("github.com/01-edu/z01")

func main() {
	for r := 'a'; r <= 'z'; r++ {
		// Print the rune and check for any errors
		if err := z01.PrintRune(r); err != nil {
			// Handle the error if needed
			// In this case, we'll just ignore the error and exit the program
			return
		}
	}
}
