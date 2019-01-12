package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
)

type point struct {
	x int
	y int
}

var grid [][]int
var minX, minY, maxX, maxY int
var width, height int

func initGrid(min_x int, min_y int, max_x int, max_y int) {
	minX = min_x
	minY = min_y
	maxX = max_x
	maxY = max_y
	width = maxX - minX + 1
	height := maxY - minY + 1
	grid = make([][]int, width)
	var i int
	for i = 0; i < width; i++ {
		grid[i] = make([]int, height)
	}
}

func setPoint(x int, y int, pointNum int) {
	grid[x-minX][y-minY] = pointNum
}

func getPoint(x int, y int) int {
	return grid[x-minX][y-minY]
}

func getEdgePoints() map[int]int {
	result := make(map[int]int)

	// "Top" edge
	for i := minX; i <= maxX; i++ {
		pointNum := getPoint(i, minY)
		result[pointNum] = 1
	}

	// "Bottom" edge
	for i := minX; i <= maxX; i++ {
		pointNum := getPoint(i, maxY)
		result[pointNum] = 1
	}

	// "Left" edge
	for i := minY; i <= maxY; i++ {
		pointNum := getPoint(minX, i)
		result[pointNum] = 1
	}

	// "Right" edge
	for i := minY; i <= maxY; i++ {
		pointNum := getPoint(maxX, i)
		result[pointNum] = 1
	}

	return result
}

func getArea(pointNum int) int {
	area := 0
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if getPoint(i, j) == pointNum {
				area++
			}
		}
	}

	return area
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

	// Read all points from the input file.
	var points []point
	var min_x int = math.MaxInt32
	var min_y int = math.MaxInt32
	var max_x int = math.MinInt32
	var max_y int = math.MinInt32
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		message := scanner.Text()
		var x, y int
		fmt.Sscanf(message, "%d, %d", &x, &y)
		p := point{x, y}
		points = append(points, p)
		if x < min_x {
			min_x = x
		}
		if y < min_y {
			min_y = y
		}
		if x > max_x {
			max_x = x
		}
		if y > max_y {
			max_y = y
		}
	}

	initGrid(min_x, min_y, max_x, max_y)

	for i := min_x; i <= max_x; i++ {
		for j := min_y; j <= max_y; j++ {
			minDistance := math.MaxInt32
			closestPoint := math.MaxInt32
			for index, p := range points {
				currentDistance := int(math.Abs(float64(p.x-i))) + int(math.Abs(float64(p.y-j)))
				if currentDistance < minDistance {
					minDistance = currentDistance
					closestPoint = index
				} else if currentDistance == minDistance {
					closestPoint = math.MaxInt32
				}
			}

			setPoint(i, j, closestPoint)
		}
	}

	// Get a list of all points that have "infinite" area
	edgePoints := getEdgePoints()

	// For each point NOT in edgePoints, calculate its area.
	maxArea := 0
	for index, _ := range points {
		_, contains := edgePoints[index]
		if !contains {
			area := getArea(index)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Printf("Max Area (including point): %d\n", maxArea)
	fmt.Printf("Max Area (excluding point): %d\n", maxArea-1)
}
