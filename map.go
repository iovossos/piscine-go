package piscine

// Map applies the function f to each element of the slice a and returns a slice of bools.
func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, elem := range a {
		result[i] = f(elem)
	}
	return result
}
