package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	// Add the last word if it exists
	if word != "" {
		result = append(result, word)
	}

	return result
}
