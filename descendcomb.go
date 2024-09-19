package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for i := 99; i >= 10; i-- {
		for j := i - 1; j >= 10; j-- {
			printTwoDigits(i)
			z01.PrintRune(' ')
			printTwoDigits(j)
			if !(i == 11 && j == 10) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func printTwoDigits(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}
