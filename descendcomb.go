package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			printTwoDigits(i)
			z01.PrintRune(' ')
			printTwoDigits(j)
			if i != 1 || j != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func printTwoDigits(n int) {
	if n < 10 {
		z01.PrintRune('0')
	}
	printNumber(n)
}

func printNumber(n int) {
	if n >= 10 {
		z01.PrintRune(rune(n/10 + '0'))
	}
	z01.PrintRune(rune(n%10 + '0'))
}
