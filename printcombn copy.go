package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	generateCombination(n, 0, []rune{})
	z01.PrintRune('\n')
}

func generateCombination(n, start int, current []rune) {
	if len(current) == n {
		for _, digit := range current {
			z01.PrintRune(digit)
		}
		if current[0] != rune('0'+10-n) { // Avoid printing comma after the last combination
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= 9; i++ {
		generateCombination(n, i+1, append(current, rune(i+'0')))
	}
}
