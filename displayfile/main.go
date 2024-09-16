package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if no arguments are passed
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}

	// Check if too many arguments are passed
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	// Open the file
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file content and print it
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println(err)
		return
	}
}
