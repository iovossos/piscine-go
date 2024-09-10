package piscine

func ToLower(s string) string {
	result := []rune(s) // Convert string to a slice of runes to handle multi-byte characters

	for i, r := range result {
		if r >= 'A' && r <= 'Z' {
			result[i] = r + ('a' - 'A') // Convert uppercase to lowercase
		}
	}

	return string(result) // Convert rune slice back to string
}
