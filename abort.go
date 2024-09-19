package piscine

func Abort(a, b, c, d, e int) int {
	mynumbers := []int{a, b, c, d, e}

	nBubblesort(mynumbers)
	return mynumbers[2]
}

func nBubblesort(mynumbers []int) []int {
	n := len(mynumbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if mynumbers[j] > mynumbers[j+1] {
				mynumbers[j], mynumbers[j+1] = mynumbers[j+1], mynumbers[j]
			}
		}
	}
	return mynumbers
}
