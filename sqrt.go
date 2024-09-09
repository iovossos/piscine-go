package piscine

func Sqrt(nb int) int {
	// Check for negative input
	if nb < 0 {
		return 0
	}

	// Special case for 0
	if nb == 0 {
		return 0
	}

	// Iterate to find the integer square root
	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}

	return 0
}
