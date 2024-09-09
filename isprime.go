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
