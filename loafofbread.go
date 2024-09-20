package piscine

func LoafOfBread(str string) string {
	var result string
	var count int // To count non-space characters

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			result += " " // Add space to the result but don't count it
			continue
		}

		if count < 5 {
			result += string(str[i]) // Add non-space characters to the result
			count++
		} else if count == 5 {
			// Skip the 6th character (i.e., reset count) and do not add it to the result
			count = 0
			continue
		}
	}

	// If there are less than 5 non-space characters, return "Invalid Output\n"
	if count < 5 && len(result) == 0 {
		return "Invalid Output\n"
	}

	return result + "\n"
}
