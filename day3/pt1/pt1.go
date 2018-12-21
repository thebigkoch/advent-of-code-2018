package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/// Parse a line with the following format:
/// #1 @ 179,662: 16x27
/// 1 == the area number
/// 179 = starting x position
/// 662 = starting y position
/// 16 = width
/// 27 = height
func parseLine(line string) (int, int, int, int) {
	var areaNum int = 0
	var x int = 0
	var y int = 0
	var width int = 0
	var height int = 0

	_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &areaNum, &x, &y, &width, &height)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return x, y, width, height
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
	fabric := [1000][1000]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		currentLine := scanner.Text()
		x, y, width, height := parseLine(currentLine)
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				fabric[x+i][y+j] += 1
			}
		}
	}

	var total int = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				total++
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
}
