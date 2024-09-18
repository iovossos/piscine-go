package main

import (
	"fmt"
	"os"
)

func checkFile(e error, name string) bool {
	if e != nil { // check if there is an error opening file
		fmt.Printf("open %s: no such file or directory\n", name)
		return false
	}
	return true // continue if no error
}

func main() {
	if len(os.Args) < 4 { // check if there are correct amount of arguments
		fmt.Println()
		os.Exit(1) // end program
	}

	numStr := os.Args[2] // store the index number
	num := 0
	for _, r := range numStr { // convert the string to an int and store in num
		num = num*10 + int(r-'0')
	}

	errExit := false
	firstFile := true
	prevHadError := false

	for _, f := range os.Args[3:] { // loop through given files
		file, err := os.Open(f) // open file with error handling
		if checkFile(err, f) {  // check error and continue if it is valid
			if !firstFile { // check if it is the first file
				if prevHadError { // check if the previous file had an error
					fmt.Printf("\n")
				} else {
					fmt.Printf("\n")
				}
			}
			fmt.Printf("==> %s <==\n", f)

			data, _ := os.ReadFile(f) // store data of the opened file
			start := len(data) - num  // get the starting index of where we want to read the file
			if start < 0 {            // check if start is negative, if it is just set the starting index to 0
				start = 0
			}
			fmt.Printf("%s", data[start:]) // print the data from the starting index till the end

			file.Close()         // close file once we are done
			prevHadError = false // no error, so we set this to false
		} else { // error has been found
			errExit = true      // bool to let us know there was an error
			prevHadError = true // let us know the previous file had an error for the next pass
		}
		firstFile = false // first file got tried so now we set the variable to false
	}

	if errExit { // check if there was an error during opening files
		os.Exit(1)
	}
}
