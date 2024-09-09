package piscine

func Fibonacci(index int) int {
	// Handle negative index
	if index < 0 {
		return -1
	}
	// Base cases
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	// Recursive case
	return Fibonacci(index-1) + Fibonacci(index-2)
}
