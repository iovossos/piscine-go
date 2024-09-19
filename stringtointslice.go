package piscine

func StringToIntSlice(str string) []int {
	var intSlice []int
	for _, char := range str {
		intSlice = append(intSlice, int(char)) // Append each rune's integer value
	}
	return intSlice
}
