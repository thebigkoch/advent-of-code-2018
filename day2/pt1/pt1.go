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
	var doubles int = 0
	var triples int = 0

	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		m := make(map[rune]int)
		currentLine := scanner.Text()

		// Iterate over the characters in the current line.
		for _, char := range currentLine {
			m[char] += 1
		}

		var doublesCounted bool = false
		var triplesCounted bool = false
		for _, v := range m {
			if (v == 2) && (!doublesCounted) {
				doubles++
				doublesCounted = true
			}
			if (v == 3) && (!triplesCounted) {
				triples++
				triplesCounted = true
			}
		}
	}

	fmt.Printf("Doubles: %d\n", doubles)
	fmt.Printf("Triples: %d\n", triples)
	fmt.Printf("d * t:   %d\n", doubles*triples)
}
