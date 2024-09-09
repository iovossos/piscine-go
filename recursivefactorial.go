package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	// Detect potential overflow by checking if (nb * factorial(nb-1)) overflows
	nextFactorial := RecursiveFactorial(nb - 1)
	if nextFactorial == 0 || nb > (1<<31-1)/nextFactorial { // check for overflow
		return 0
	}
	return nb * nextFactorial
}
