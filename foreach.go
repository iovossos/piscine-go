package piscine

// ForEach applies the function f to each element of the slice a.
func ForEach(f func(int), a []int) {
	for _, elem := range a {
		f(elem)
	}
}
