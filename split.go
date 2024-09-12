package piscine

func Split(s, sep string) []string {
	var result []string
	var temp string
	sepLen := len(sep)
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		// Check if the current substring matches the separator
		if i+sepLen <= sLen && s[i:i+sepLen] == sep {
			result = append(result, temp)
			temp = ""
			i += sepLen - 1 // Skip the separator
		} else {
			temp += string(s[i])
		}
	}

	// Add the remaining part after the last separator
	if temp != "" {
		result = append(result, temp)
	}

	return result
}
