package piscine

func IterativePower(nb int, power int) int {
	// Return 0 for negative powers
	if power < 0 {
		return 0
	}

	// Initialize result to 1 (the neutral element for multiplication)
	result := 1

	// Calculate nb to the power of power iteratively
	for i := 0; i < power; i++ {
		result *= nb
	}

	return result
}
