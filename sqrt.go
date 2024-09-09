package piscine

import (
	"math"
)

func Sqrt(nb int) int {
	// Check for negative input
	if nb < 0 {
		return 0
	}

	// Calculate the square root and round it down to the nearest integer
	root := int(math.Sqrt(float64(nb)))

	// Check if squaring the root gives back the original number
	if root*root == nb {
		return root
	}

	return 0
}
