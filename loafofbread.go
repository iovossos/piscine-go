package piscine

func LoafOfBread(str string) string {
	var result string
	var count int // To count characters excluding spaces

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' { // Ignore spaces
			continue
		}

		if count < 5 {
			result += string(str[i]) // Add characters to the result until 5 chars are collected
			count++
		} else if count == 5 {
			// Skip the 6th character
			count = 0
			continue
		}
	}

	// If the string has less than 5 characters, return "Invalid Output\n"
	if len(result) < 5 {
		return "Invalid Output\n"
	}

	// Return the formatted result with a newline at the end
	return result + "\n"
}
