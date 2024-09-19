package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			runes[i] = 'a' + (char-'a'+14)%26
		} else if char >= 'A' && char <= 'Z' {
			runes[i] = 'A' + (char-'A'+14)%26
		}
	}
	return string(runes)
}
