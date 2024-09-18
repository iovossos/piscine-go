package piscine

// AdvancedSortWordArr sorts a slice of strings based on the comparison function f.
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if f(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
