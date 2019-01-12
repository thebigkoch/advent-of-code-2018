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
func parseLine(line string) (int, int, int, int, int) {
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

	return areaNum, x, y, width, height
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
		_, x, y, width, height := parseLine(currentLine)
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				fabric[x+i][y+j] += 1
			}
		}
	}

	// Now individually check each of the areas to see if something else overlaps
	f, err = os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	defer f.Close()
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		currentLine := scanner.Text()
		area, x, y, width, height := parseLine(currentLine)
		overlaps := false
		for i := 0; i < width && !overlaps; i++ {
			for j := 0; j < height && !overlaps; j++ {
				if fabric[x+i][y+j] > 1 {
					overlaps = true
				}
			}
		}
		if !overlaps {
			fmt.Printf("There is no overlap on area #%d", area)
		}
	}
}
