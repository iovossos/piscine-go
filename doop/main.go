package main

import (
	"math"
	"os"
)

func atoi(s string) (int64, bool) {
	result := int64(0)
	sign := int64(1)
	i := 0

	if len(s) == 0 {
		return 0, false
	}

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		result = result*10 + int64(s[i]-'0')

		// Check for overflow
		if result > math.MaxInt64 {
			return 0, false
		}
	}

	return result * sign, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	// Parse the first and third arguments manually using atoi
	a, okA := atoi(os.Args[1])
	b, okB := atoi(os.Args[3])
	if !okA || !okB {
		return
	}

	// Get the operator (second argument)
	operator := os.Args[2]

	switch operator {
	case "+":
		// Check for overflow
		if (b > 0 && a > math.MaxInt64-b) || (b < 0 && a < math.MinInt64-b) {
			return
		}
		printNbr(a + b)

	case "-":
		// Check for overflow
		if (b < 0 && a > math.MaxInt64+b) || (b > 0 && a < math.MinInt64+b) {
			return
		}
		printNbr(a - b)

	case "*":
		// Check for overflow
		if a != 0 && b != 0 && (a > math.MaxInt64/b || a < math.MinInt64/b) {
			return
		}
		printNbr(a * b)

	case "/":
		if b == 0 {
			printStr("No division by 0\n")
			return
		}
		printNbr(a / b)

	case "%":
		if b == 0 {
			printStr("No modulo by 0\n")
			return
		}
		printNbr(a % b)

	default:
		// If the operator is invalid, print nothing
		return
	}
}

func printNbr(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte("0\n"))
		return
	}

	if n < 0 {
		os.Stdout.Write([]byte("-"))
		n = -n
	}

	digits := []byte{}
	for n > 0 {
		digits = append([]byte{byte(n%10 + '0')}, digits...)
		n /= 10
	}

	os.Stdout.Write(digits)
	os.Stdout.Write([]byte("\n"))
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
}
