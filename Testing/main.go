package main

import (
	"fmt"
)

func main() {
	PrintMenu()
}

func PrintMenu() {
	fmt.Println("╔═══════════════════════════╗")
	fmt.Println("║                           ║")
	fmt.Println("║                           ║")
	fmt.Print("║          ")
	fmt.Printf("\033[1m%s\033[0m", "Sudoku!")
	fmt.Println("          ║")
	fmt.Println("║                           ║")
	fmt.Println("║     Select an option:     ║")
	fmt.Println("╚═══════════════════════════╝")
}
