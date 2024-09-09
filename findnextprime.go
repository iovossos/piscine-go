package piscine

func IsPrime(nb int) bool {
	// Handle special cases
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true // 2 is the only even prime number
	}
	if nb%2 == 0 {
		return false // Exclude even numbers greater than 2
	}

	// Check for factors from 3 up to the square root of nb
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}

	return true
}

func FindNextPrime(nb int) int {
	// Start checking from nb and keep incrementing until a prime is found
	for nb >= 0 {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
	return 0 // This line will never be reached because nb will always be >= 0
}
