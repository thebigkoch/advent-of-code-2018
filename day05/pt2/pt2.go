package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func removeLetter(input string, letter rune) string {
	result := strings.Replace(input, strings.ToLower(string(letter)), "", -1)
	result = strings.Replace(result, strings.ToUpper(string(letter)), "", -1)
	return result
}

func eliminateOpposites(input string) string {
	skippedLetters := 1
	var result string
	var upperToLowerDiff int = 'A' - 'a'
	for skippedLetters > 0 {
		result = ""
		skippedLetters = 0
		for i := 0; i < len(input); i++ {
			if i == len(input)-1 {
				result += string(input[i])
			} else if (int(input[i])-int(input[i+1]) == upperToLowerDiff) ||
				(int(input[i+1])-int(input[i]) == upperToLowerDiff) {
				i++
				skippedLetters++
			} else {
				result += string(input[i])
			}
		}
		input = result
	}

	return result
}

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
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		message = scanner.Text()
	}

	smallestResult := len(message)
	for char := 'a'; char <= 'z'; char++ {
		messageWithoutChar := removeLetter(message, char)
		messageWithoutOpposites := eliminateOpposites(messageWithoutChar)
		fmt.Printf("len(message) without %c = %d\n", char, len(messageWithoutOpposites))
		if len(messageWithoutOpposites) < smallestResult {
			smallestResult = len(messageWithoutOpposites)
		}
	}

	fmt.Printf("Smallest result = %d\n", smallestResult)
}
