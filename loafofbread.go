package piscine

func LoafOfBread(str string) string {
	// If the string is less than 5 characters, return "Invalid Output\n"
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	count := 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' { // Ignore spaces
			continue
		}

		if count < 5 {
			result += string(str[i]) // Add characters to the result until 5 chars are collected
			count++
		} else if count == 5 {
			// Skip the 6th character and add a space
			count = 0
			result += " "
			continue
		}
	}

	// Remove the trailing space if present
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
