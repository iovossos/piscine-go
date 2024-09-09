package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		prevResult := result
		result *= i
		// Check if overflow occurred by seeing if result is smaller than the previous value
		if result/i != prevResult {
			return 0 // Overflow detected
		}
	}
	return result
}
