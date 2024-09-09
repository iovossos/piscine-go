package piscine

func RecursiveFactorial(nb int64) int64 {
	// Return 0 for negative numbers and numbers too large for factorial
	if nb < 0 || nb > 20 { // 20! is the largest factorial that fits in an int64
		return 0
	}
	return factorialHelper(nb, 1)
}

// Helper function to carry the accumulated result
func factorialHelper(nb int64, accumulator int64) int64 {
	if nb == 0 || nb == 1 {
		return accumulator
	}
	// Overflow check for 64-bit integers
	if nb > (1<<63-1)/accumulator {
		return 0 // Overflow detected
	}
	return factorialHelper(nb-1, accumulator*nb)
}
