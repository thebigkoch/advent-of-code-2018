package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type pointDef struct {
	initialX int
	initialY int
	dX       int
	dY       int
}

func calculatePoints(pointdefs []pointDef, time int) []point {
	var points []point

	for _, pd := range pointdefs {
		p := point{pd.initialX + pd.dX*time, pd.initialY + pd.dY*time}
		points = append(points, p)
	}

	return points
}

func isValidMessage(points []point) bool {

	// We consider the message "valid" if each point has a neighboring point.
	for i1, p1 := range points {
		hasNeighbor := false
		for i2, p2 := range points {
			if (i1 != i2) && !hasNeighbor {
				diffX := p1.x - p2.x
				diffY := p1.y - p2.y
				if (math.Abs(float64(diffX)) <= 1.0) && (math.Abs(float64(diffY)) <= 1.0) {
					hasNeighbor = true
				}
			}
		}

		if !hasNeighbor {
			return false
		}
	}
	return true
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

	// Read input file.
	var pointdefs []pointDef
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		initialX, _ := strconv.Atoi(strings.TrimSpace(line[10:16]))
		initialY, _ := strconv.Atoi(strings.TrimSpace(line[17:24]))
		dX, _ := strconv.Atoi(strings.TrimSpace(line[36:38]))
		dY, _ := strconv.Atoi(strings.TrimSpace(line[39:42]))

		pointdef := pointDef{initialX, initialY, dX, dY}
		pointdefs = append(pointdefs, pointdef)
	}

	var time int = 0
	var points []point
	for {
		// Calculate the position of points for the current time.
		points = calculatePoints(pointdefs, time)
		if isValidMessage(points) {
			break
		}
		time++
	}

	// Calculate the bounding box of the message.
	var minX int = points[0].x
	var minY int = points[0].y
	var maxX int = points[0].x
	var maxY int = points[0].y

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// Now print the output
	fmt.Printf("Message found at time=%d in bounding box {minX=%d, minY=%d, maxX=%d, maxY=%d}\n", time, minX, minY, maxX, maxY)
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			containsPoint := false
			for _, p := range points {
				if (p.y == i) && (p.x == j) {
					containsPoint = true
				}
			}

			if containsPoint {
				fmt.Printf("X")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
