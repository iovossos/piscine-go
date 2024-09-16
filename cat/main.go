package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printError(err error) {
	errorMessage := "ERROR: " + err.Error() + "\n"
	for _, r := range errorMessage {
		z01.PrintRune(r)
	}
}

func printContent(reader io.Reader) {
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			printError(err)
			return
		}
		for _, r := range buffer[:n] {
			z01.PrintRune(rune(r))
		}
	}
}

func main() {
	args := os.Args[1:]

	// If no arguments are passed, read from stdin
	if len(args) == 0 {
		printContent(os.Stdin)
		return
	}

	// Iterate through the provided arguments (files)
	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			printError(err)
			continue
		}
		printContent(file)
		file.Close()
	}
}
