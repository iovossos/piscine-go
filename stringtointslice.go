package piscine

func StringToIntSlice(str string) []int {
	pinakas := make([]int, len(str))
	for i, char := range str {
		pinakas[i] = int(char)
	}
	return pinakas
}
