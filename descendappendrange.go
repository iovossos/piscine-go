package piscine

func DescendAppendRange(max, in int) []int {
	if max >= min {
		return nil // Return nil slice if min is greater than or equal to max
	}

	// Initialize an empty slice (without using make)
	var result []int

	// Append values from min to max - 1
	for i := max; i < min; i++ {
		result = append(result, i)
	}

	return result
}
