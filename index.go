package piscine

func Index(s string, toFind string) int {
	runesS := []rune(s)
	runesToFind := []rune(toFind)

	lenS := len(runesS)
	lenToFind := len(runesToFind)

	if lenToFind == 0 || lenToFind > lenS {
		return -1
	}

	for i := 0; i <= lenToFind; i++ {
		match := true
		for j := 0; j < lenToFind; j++ {
			if runesS[i+j] != runesToFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
