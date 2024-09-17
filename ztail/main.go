package main

import (
	"fmt"
	"os"
)

func printError(fileName string, err error) {
	// Print a clean error message without duplication
	fmt.Printf("open %s: %s\n", fileName, err.Error())
}

func stringToInt(s string) (int, error) {
	result := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("Invalid number: %s", s)
		}
		result = result*10 + int(char-'0')
	}
	return result, nil
}

func tailFile(fileName string, count int) bool {
	file, err := os.Open(fileName)
	if err != nil {
		printError(fileName, err)
		return false
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		printError(fileName, err)
		return false
	}

	// Calculate where to start reading
	size := fileInfo.Size()
	offset := size - int64(count)
	if offset < 0 {
		offset = 0
	}

	// Move the file pointer to the start of the last `count` characters
	_, err = file.Seek(offset, 0)
	if err != nil {
		printError(fileName, err)
		return false
	}

	// Read the last `count` characters and print them
	buffer := make([]byte, count)
	n, err := file.Read(buffer)
	if err != nil {
		printError(fileName, err)
		return false
	}

	// Print the content
	fmt.Printf("%s", string(buffer[:n]))
	return true
}

func main() {
	args := os.Args[1:]

	// Check if there are enough arguments
	if len(args) < 3 {
		fmt.Printf("Usage: go run . -c <number> <file1> <file2> ...\n")
		os.Exit(1)
	}

	// Parse the -c option
	if args[0] != "-c" {
		fmt.Printf("Invalid option. Only '-c' is allowed.\n")
		os.Exit(1)
	}

	// Get the number of characters to display
	count, err := stringToInt(args[1])
	if err != nil || count < 0 {
		fmt.Printf("Invalid number of characters: %s\n", args[1])
		os.Exit(1)
	}

	// Process each file
	exitStatus := 0
	firstFileProcessed := false
	for i := 2; i < len(args); i++ {
		if !tailFile(args[i], count) {
			// Mark exit status if there was an error
			exitStatus = 1
		} else {
			// Print the filename header for valid files
			if firstFileProcessed {
				fmt.Printf("\n==> %s <==\n", args[i])
			} else {
				fmt.Printf("==> %s <==\n", args[i])
				firstFileProcessed = true
			}
		}
	}

	// Exit with the appropriate status
	if exitStatus != 0 {
		os.Exit(1)
	}
}
