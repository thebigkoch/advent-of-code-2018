package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func differingLetters(s1 string, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}

	var result int = 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			result++
		}
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

	// Read all strings into an array.
	var strings []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		currentLine := scanner.Text()
		strings = append(strings, currentLine)
	}

	for i, s1 := range strings {
		for j, s2 := range strings {
			if i != j {
				differingRunes := differingLetters(s1, s2)
				if differingRunes == 1 {
					fmt.Printf("String 1: %s\n", s1)
					fmt.Printf("String 2: %s\n\n", s2)
				}
			}
		}
	}
}
