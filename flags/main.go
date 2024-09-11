package main

import (
	"fmt"
	"os"
)

func helpPrnt() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func sortStr(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		helpPrnt()
		return
	}

	var insertString string
	var mainString string
	order := false

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "-h" || arg == "--help" {
			helpPrnt()
			return
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insertString = arg[3:]
		} else if len(arg) > 9 && arg[:9] == "--insert=" {
			insertString = arg[9:]
		} else if arg == "-0" || arg == "--order" {
			order = true
		} else {
			mainString = arg
		}
	}
	if !(order && insertString == "") {
		mainString = mainString + insertString
	}
	if !order && insertString == "" {
		mainString = sortStr(mainString)
	}
	if order && insertString != "" {
		mainString = sortStr(mainString + insertString)
	}
	if order && insertString == "" {
		mainString = sortStr(mainString)
	}
	fmt.Println(mainString)
}
