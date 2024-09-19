package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reversedMenu := make([]string, length)

	for i := 0; i < length; i++ {
		reversedMenu[i] = menu[length-i-1]
	}

	return reversedMenu
}
