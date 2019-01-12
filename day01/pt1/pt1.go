package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

	var total int64 = 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		total += i
	}
	fmt.Printf("Total: %d\n", total)
}
