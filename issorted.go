package piscine

// IsSorted checks if the slice a is sorted based on the comparison function f.
// It handles both ascending and descending order.
func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	// Determine the sorting direction based on the first two elements
	isAscending := f(a[0], a[1]) <= 0

	for i := 1; i < len(a); i++ {
		if isAscending {
			// Check if the slice is not sorted in ascending order
			if f(a[i-1], a[i]) > 0 {
				return false
			}
		} else {
			// Check if the slice is not sorted in descending order
			if f(a[i-1], a[i]) < 0 {
				return false
			}
		}
	}

	return true
}
