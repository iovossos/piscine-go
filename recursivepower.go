package piscine

func RecursivePower(nb int, power int) int {
	// Handle negative powers
	if power < 0 {
		return 0
	}
	// Base case for power 0
	if power == 0 {
		return 1
	}
	// Recursive case
	return nb * RecursivePower(nb, power-1)
}
