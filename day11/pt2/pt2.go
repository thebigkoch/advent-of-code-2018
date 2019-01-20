package main

import "fmt"

const MIN_X int = 1
const MIN_Y int = 1
const MAX_X int = 300
const MAX_Y int = 300
const MIN_SIZE int = 1
const MAX_SIZE int = 300

type point struct {
	x int
	y int
}

// Calculate the power of a specific point.
func calculatePowerOfPoint(x int, y int, serialNumber int) int {

	rackId := x + 10
	powerLevel := rackId * y
	powerLevel += serialNumber
	powerLevel *= rackId

	// Get the hundreds digit of powerLevel.
	powerLevel = int(powerLevel / 100)
	powerLevel = powerLevel % 10
	powerLevel -= 5

	return powerLevel
}

// Calculate the power of an area with top-level coord (x,y).
func calculatePowerOfArea(x int, y int, previousPower int, size int, serialNumber int) int {
	powerOfArea := previousPower
	for i := 0; i < size; i++ {
		j := size - 1
		if ((x + i) > MIN_X) &&
			((x + i) < MAX_X) &&
			((y + j) > MIN_Y) &&
			((y + j) < MAX_Y) {
			powerOfArea += calculatePowerOfPoint(x+i, y+j, serialNumber)
		}
	}

	for j := 0; j < size-1; j++ {
		i := size - 1
		if ((x + i) > MIN_X) &&
			((x + i) < MAX_X) &&
			((y + j) > MIN_Y) &&
			((y + j) < MAX_Y) {
			powerOfArea += calculatePowerOfPoint(x+i, y+j, serialNumber)
		}
	}

	return powerOfArea
}

func main() {
	// My puzzle input is 5093.
	serialNumber := 5093

	// Calculate the power of each point.
	mapPowerOfPoint := make(map[point]int)

	for x := MIN_X; x <= MAX_X; x++ {
		for y := MIN_Y; y <= MAX_Y; y++ {
			p := point{x, y}
			mapPowerOfPoint[p] = calculatePowerOfPoint(x, y, serialNumber)
		}
	}

	maxPower := 0
	maxX := 0
	maxY := 0
	maxSize := 0

	// Calculate the power of each grid area.
	for x := MIN_X; x <= MAX_X; x++ {
		fmt.Printf("x = %d\n", x)
		for y := MIN_Y; y <= MAX_Y; y++ {
			previousPower := 0
			for size := MIN_SIZE; size <= (MAX_SIZE - x + 1); size++ {
				currentPower := calculatePowerOfArea(x, y, previousPower, size, serialNumber)
				if currentPower > maxPower {
					maxPower = currentPower
					maxX = x
					maxY = y
					maxSize = size
				}
				previousPower = currentPower
			}
		}
	}
	fmt.Printf("Max power = %d at (%d, %d, %d)\n", maxPower, maxX, maxY, maxSize)
}
