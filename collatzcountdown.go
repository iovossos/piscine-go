package piscine

// var count int = 0

// func CollatzCountdown(start int) int {
// 	if start == 1 {
// 		return count
// 	} else if start%2 == 0 {
// 		count++
// 		CollatzCountdown((start / 2))
// 	} else if start%2 == 1 {
// 		count++
// 		CollatzCountdown((start * 3) + 1)
// 	}
// 	return count
// }

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	steps := 0

	for start != 1 {

		if start%2 == 0 {
			start /= 2
		} else {
			start = 3*start + 1
		}

		steps++

	}

	return steps
}
