package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	var result string
	for i, str := range strs {
		if i > 0 {
			result += sep
		}
		result += str
	}
	return result
}
