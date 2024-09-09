package piscine

// Helper function to determine if a number is prime
func isPrimeNumber(nb int) bool {
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

// Function to find the next prime number greater than or equal to nb
func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2 // The smallest prime number is 2
	}
	if nb == 2 {
		return 2 // If nb is 2, it is the prime itself
	}
	// Start checking from nb or nb+1 if nb is even
	if nb%2 == 0 {
		nb++
	}
	// Increment and check each odd number for primality
	for ; !isPrimeNumber(nb); nb += 2 {
	}
	return nb
}
