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

func tailFile(fileName string, count int) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return err
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
		return err
	}

	// Read the last `count` characters
	buffer := make([]byte, count)
	n, err := file.Read(buffer)
	if err != nil {
		return err
	}

	// Print the content of the file
	fmt.Printf("%s", string(buffer[:n]))
	return nil
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
	for i := 2; i < len(args); i++ {
		// Try to tail the file
		err := tailFile(args[i], count)
		if err != nil {
			// Print error if file cannot be processed
			printError(args[i], err)
			exitStatus = 1
		} else {
			// Print the filename header for valid files before content
			if i != 2 {
				fmt.Printf("\n==> %s <==\n", args[i])
			} else {
				fmt.Printf("==> %s <==\n", args[i])
			}
		}
	}

	// Exit with the appropriate status
	if exitStatus != 0 {
		os.Exit(1)
	}
}
