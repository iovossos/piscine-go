package piscine

func StringToIntSlice(str string) []int {
	var pinakas []int
	for i, char := range str {
		pinakas[i] = int(char)
	}
	return pinakas
}
