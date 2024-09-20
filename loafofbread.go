package piscine

func LoafOfBread(str string) string {
	var result string
	var wordCount int // To count characters excluding the skipped ones
	var charCount int // To count all characters, including spaces

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' { // Include spaces in the result
			result += " "
			continue
		}

		if wordCount < 5 {
			result += string(str[i]) // Add characters to the result until 5 chars are collected
			wordCount++
		} else if wordCount == 5 {
			// Skip the 6th character and reset the word count
			wordCount = 0
			continue
		}

		charCount++
	}

	// If the string has less than 5 characters, return "Invalid Output\n"
	if charCount < 5 {
		return "Invalid Output\n"
	}

	// Return the formatted result with a newline at the end
	return result + "\n"
}
