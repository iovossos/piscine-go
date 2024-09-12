package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil // Return nil slice if min is greater than or equal to max
	}

	// Initialize an empty slice (without using make)
	var result []int

	// Append values from min to max - 1
	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}
