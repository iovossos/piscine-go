package piscine

func RecursiveFactorial(nb int) int64 {
	// Immediately return 0 for negative inputs
	if nb < 0 {
		return 0
	}
	return factorialHelper(nb, 1)
}

// Helper function to carry the accumulated result
func factorialHelper(nb int, accumulator int64) int64 {
	if nb == 0 || nb == 1 {
		return accumulator
	}
	// Overflow check for 64-bit integers
	if nb > int((1<<63-1)/accumulator) {
		return 0 // Overflow detected
	}
	return factorialHelper(nb-1, accumulator*int64(nb))
}
