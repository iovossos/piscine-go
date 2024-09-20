package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}

	// Check if divisible by both 2 and 3
	if n%2 == 0 && n%3 == 0 {
		return "rock and roll\n"
	}

	// Check if divisible by 2
	if n%2 == 0 {
		return "rock\n"
	}

	// Check if divisible by 3
	if n%3 == 0 {
		return "roll\n"
	}

	// If not divisible by 2 or 3
	return "error: non divisible\n"
}
