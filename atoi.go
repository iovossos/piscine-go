package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	startIndex := 0

	// Handle optional sign
	if len(s) > 0 {
		if s[0] == '-' {
			sign = -1
			startIndex = 1
		} else if s[0] == '+' {
			startIndex = 1
		}
	}

	// Convert the digits to an integer
	for i := startIndex; i < len(s); i++ {
		char := s[i]
		if char < '0' || char > '9' {
			// Return 0 if any character is not a digit
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result * sign
}
