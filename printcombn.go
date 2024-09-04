package piscine

import (
	"github.com/01-edu/z01"
)

// PrintCombN prints all unique combinations of n different digits in ascending order.
func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	// Start generating combinations
	generateComb(0, n, 0)
	z01.PrintRune('\n') // Print a newline at the end
}

// generateComb is a recursive helper function to generate and print combinations
func generateComb(start, n, length int) {
	if length == n {
		printCombination()
		return
	}
	for i := start; i <= 9; i++ {
		combo[length] = i
		generateComb(i+1, n, length+1)
	}
}

// printCombination prints the current combination stored in combo
func printCombination() {
	for _, digit := range combo {
		z01.PrintRune(rune('0' + digit))
	}
	if !isLastCombination() {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

// isLastCombination checks if the current combination is the last one
func isLastCombination() bool {
	for i := 0; i < n-1; i++ {
		if combo[i] != combo[i+1]-1 {
			return false
		}
	}
	return true
}

// Global variable to store current combination
var combo [10]int
var n int
