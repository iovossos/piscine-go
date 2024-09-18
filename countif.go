package piscine

// CountIf returns the number of elements in the slice tab for which the function f returns true.
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, elem := range tab {
		if f(elem) {
			count++
		}
	}
	return count
}
