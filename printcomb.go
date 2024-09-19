package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb1() {
	// Loop through all possible combinations of three different digits
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)

				// Check if it is not the last combination
				if !(i == '7' && j == '8' && k == '9') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n') // Print a newline at the end
}
