package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	// Prevent overflow for 32-bit integers by checking if multiplication would exceed the limit
	nextFactorial := RecursiveFactorial(nb - 1)
	if nextFactorial == 0 || nb > (1<<31-1)/nextFactorial {
		return 0
	}
	return nb * nextFactorial
}
