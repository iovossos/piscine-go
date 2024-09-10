package piscine

func Index(s string, toFind string) int {
	// Convert both strings to rune slices for proper handling of UTF-8 characters.
	runesS := []rune(s)
	runesToFind := []rune(toFind)

	lenS := len(runesS)
	lenToFind := len(runesToFind)

	// If toFind is empty, return 0 (empty string is considered to be found at the start).
	if lenToFind == 0 {
		return 0
	}

	// If toFind is longer than s, return -1.
	if lenToFind > lenS {
		return -1
	}

	// Loop through the string and look for the substring.
	for i := 0; i <= lenS-lenToFind; i++ {
		match := true
		// Check if substring matches starting at index i.
		for j := 0; j < lenToFind; j++ {
			if runesS[i+j] != runesToFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	// Return -1 if the substring is not found.
	return -1
}
