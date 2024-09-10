package piscine

func Compare(a, b string) int {
	runesA := []rune(a)
	runesB := []rune(b)

	minLen := len(runesA)
	if len(runesB) < minLen {
		minLen = len(runesB)
	}

	for i := 0; i < minLen; i++ {
		if runesA[i] != runesB[i] {
			if runesA[i] < runesB[i] {
				return -1
			}
			return 1
		}
	}

	if len(runesA) == len(runesB) {
		return 0
	} else if len(runesA) < len(runesB) {
		return -1
	}
	return 1
}
