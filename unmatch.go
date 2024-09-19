package piscine

func Unmatch(a []int) int {
	frequency := make(map[int]int)
	for _, num := range a {
		frequency[num]++
	}
	for _, num := range a {
		if frequency[num]%2 != 0 {
			return num
		}
	}

	return -1
}
