package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get the current directory.
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Open the input file.
	filePath := filepath.Join(currentDir, "input.txt")
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	// Read the line into a temporary variable.
	var message string
	var upperToLowerDiff int = 'A' - 'a'
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		message = scanner.Text()
	}

	skippedLetters := 1
	var newMessage string
	for skippedLetters > 0 {
		newMessage = ""
		skippedLetters = 0
		for i := 0; i < len(message); i++ {
			if i == len(message)-1 {
				newMessage += string(message[i])
			} else if (int(message[i])-int(message[i+1]) == upperToLowerDiff) ||
				(int(message[i+1])-int(message[i]) == upperToLowerDiff) {
				i++
				skippedLetters++
			} else {
				newMessage += string(message[i])
			}
		}
		message = newMessage
	}

	fmt.Printf("Final result: %s\n", message)
	fmt.Printf("Final len = %d\n", len(message))
}
