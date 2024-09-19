package main

import (
	"fmt"
	"os"
)

func main() {
	keywords := []string{"01", "galaxy", "galaxy 01"}
	for _, arg := range os.Args[1:] {
		for _, keyword := range keywords {
			if arg == keyword {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
