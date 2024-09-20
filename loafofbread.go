package piscine

func LoafOfBread(str string) string {
	var result string
	var count int // To count non-space characters

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			result += " " // Add space to the result
			continue
		}

		if count < 5 {
			result += string(str[i]) // Add characters to the result until 5 non-space chars are collected
			count++
		} else if count == 5 {
			// Skip the 6th character (reset count) but do not add it to the result
			count = 0
			continue
		}
	}

	// If the string has less than 5 non-space characters, return "Invalid Output\n"
	if len(result) < 5 {
		return "Invalid Output\n"
	}

	return result + "\n"
}
