package piscine

// Any returns true if at least one element in the string slice a returns true when passed to function f.
func Any(f func(string) bool, a []string) bool {
	for _, elem := range a {
		if f(elem) {
			return true
		}
	}
	return false
}
