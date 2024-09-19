package piscine

func StringToSlice(str string) []int {
	var pinakas []int
	for i, char := range str {
		pinakas[i] = int(char)
	}
	return pinakas
}
