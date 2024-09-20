package piscine

func LoafOfBread(str string) string {
	var result string
	var count int // To count non-space characters

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			result += " " // Add spaces to the result but don't count them
			continue
		}

		if count < 5 {
			result += string(str[i]) // Add non-space characters to the result
			count++
		} else if count == 5 {
			// Skip the 6th non-space character and reset the count
			count = 0
			continue
		}
	}

	// If there are fewer than 5 non-space characters, return "Invalid Output\n"
	nonSpaceCount := 0
	for i := 0; i < len(result); i++ {
		if result[i] != ' ' {
			nonSpaceCount++
		}
	}

	if nonSpaceCount < 5 {
		return "Invalid Output\n"
	}

	return result + "\n"
}
