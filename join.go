package piscine

func Join(strs []string, sep string) string {
	var result string
	for _, elem := range strs {
		result += elem
		result += sep
	}
	result += sep
	return result
}
