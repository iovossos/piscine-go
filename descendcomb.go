package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			printTwoDigits(i)
			z01.PrintRune(' ')
			printTwoDigits(j)
			if !(i == 1 && j == 0) { // To avoid the trailing comma and space
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}

func printTwoDigits(n int) {
	// Print the first digit (tens place)
	z01.PrintRune(rune(n/10 + '0'))
	// Print the second digit (ones place)
	z01.PrintRune(rune(n%10 + '0'))
}
