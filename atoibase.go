package piscine

func AtoiBase(s string, base string) int {
	// Define the max int64 value for overflow checks
	const MaxInt = 9223372036854775807

	// Check if the base is valid
	if !isValidBase1(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	// Map each character in base to its index for quick lookup
	baseMap := make(map[rune]int)
	for i, r := range base {
		baseMap[r] = i
	}

	// Convert the string to an integer based on the base
	for _, char := range s {
		index, exists := baseMap[char]
		if !exists { // If the character is not in the base, it's an invalid string
			return 0
		}

		// Check for potential overflow before multiplying by baseLen
		if result > (MaxInt-index)/baseLen {
			return 0 // Overflow detected, return 0
		}

		result = result*baseLen + index
	}

	return result
}

// Helper function to check base validity
func isValidBase1(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
