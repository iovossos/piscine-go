package piscine

func LoafOfBread(str string) string {
	var result string
	var count int // To count non-space characters

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			result += " " // Add space directly to the result
			continue
		}

		if count < 5 {
			result += string(str[i]) // Add non-space characters to the result
			count++
		} else if count == 5 {
			// Skip the 6th character and reset count
			count = 0
			continue
		}
	}

	// If the total number of non-space characters is less than 5, return "Invalid Output\n"
	if count > 0 && count < 5 {
		return "Invalid Output\n"
	}

	return result + "\n"
}
