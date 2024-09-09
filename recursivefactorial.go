package piscine

func RecursiveFactorial(nb int) int {
	// Upper bound check to prevent deep recursion and overflow
	if nb < 0 || nb > 12 { // 12! is the largest factorial that fits in a 32-bit int
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
