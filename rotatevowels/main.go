package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}

func mirror(runes []rune) {
	// Collect vowel positions from the entire sentence
	var vowels []int
	for i, r := range runes {
		if isVowel(r) {
			vowels = append(vowels, i)
		}
	}

	// Mirror vowels across the entire sentence
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		runes[vowels[i]], runes[vowels[j]] = runes[vowels[j]], runes[vowels[i]]
	}
}

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	// Combine all arguments into one slice of runes
	var allRunes []rune
	for i, arg := range os.Args[1:] {
		if i > 0 {
			allRunes = append(allRunes, ' ') // Add a space between arguments
		}
		allRunes = append(allRunes, []rune(arg)...)
	}

	// Mirror vowels across the whole sentence
	mirror(allRunes)

	// Print the result
	for _, r := range allRunes {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
