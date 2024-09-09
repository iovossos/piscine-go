package piscine

func RecursiveFactorial(nb int) int64 {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}

	// Prevent overflow for 64-bit integers
	nextFactorial := RecursiveFactorial(nb - 1)
	if nb > int((1<<63-1)/nextFactorial) {
		return 0 // Overflow detected
	}

	return int64(nb) * nextFactorial
}
