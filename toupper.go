package piscine

func ToUpper(s string) string {
	result := []rune(s) // Convert string to a slice of runes to handle multi-byte characters

	for i, r := range result {
		if r >= 'a' && r <= 'z' {
			result[i] = r - ('a' - 'A') // Convert lowercase to uppercase
		}
	}

	return string(result) // Convert rune slice back to string
}
