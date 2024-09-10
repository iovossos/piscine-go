package piscine

func Capitalize(s string) string {
	var result []rune
	shouldUpper := true
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			if shouldUpper {
				if ch >= 'a' && ch <= 'z' {
					ch -= 32
				}
				shouldUpper = false
			} else {
				if ch >= 'A' && ch <= 'Z' {
					ch += 32
				}
			}
			result = append(result, ch)
		} else {
			result = append(result, ch)
			if !(ch >= 'a' && ch <= 'z') && !(ch >= 'A' && ch <= 'Z') && !(ch >= '0' && ch <= '9') {
				shouldUpper = true
			}
		}
	}

	return string(result)
}
