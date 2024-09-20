package piscine

func LoafOfBread(str string) string {
	// If the string is empty, return just a newline
	if len(str) == 0 {
		return "\n"
	}

	// If the string has less than 5 characters, return "Invalid Output\n"
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	count := 0

	for i := 0; i < len(str); i++ {
		if count == 5 {
			// Skip the 6th character and reset the count
			i++
			count = 0
			// Only add a space if there are more characters left and they are not spaces
			for i < len(str) && str[i] == ' ' {
				i++ // Skip any extra spaces
			}
			if i < len(str) { // Add space only if there are more valid characters left
				result += " "
			}
			if i >= len(str) {
				break
			}
		}
		if i < len(str) && str[i] != ' ' { // Ignore spaces in the string
			result += string(str[i])
			count++
		}
	}

	// Append newline character at the end without trailing spaces
	return result + "\n"
}
