package piscine

// IsSorted checks if the slice a is sorted in ascending order based on the comparison function f.
func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 { // if previous element is greater than current, not sorted
			return false
		}
	}
	return true
}
