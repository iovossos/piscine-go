package piscine

func JumpOver(str string) string {
	count := 0
	output := ""
	for _, char := range str {
		count++
		if count%3 == 0 {
			output += string(char)
		}
	}
	output += "\n"
	return output
}
