package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
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

	messages := make(map[string]string)
	var kDates []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		currentLine := scanner.Text()
		strDateTime := currentLine[1:17]
		strMsg := currentLine[19:]
		messages[strDateTime] = strMsg
		kDates = append(kDates, strDateTime)
	}
	sort.Strings(kDates)

	for _, strDateTime := range kDates {
		fmt.Printf("Date: %s\n", strDateTime)
		fmt.Printf("Message: %s\n\n", messages[strDateTime])
	}
}
