package main

import "fmt"

const MIN_X int = 1
const MIN_Y int = 1
const MAX_X int = 300
const MAX_Y int = 300

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
func calculatePowerOfArea(x int, y int, serialNumber int) int {
	powerOfArea := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if ((x + i) > MIN_X) &&
				((x + i) < MAX_X) &&
				((y + j) > MIN_Y) &&
				((y + j) < MAX_Y) {
				powerOfArea += calculatePowerOfPoint(x+i, y+j, serialNumber)
			}
		}
	}

	return powerOfArea
}

func main() {
	// My puzzle input is 5093.
	serialNumber := 5093
	maxPower := 0
	maxX := 0
	maxY := 0

	for x := MIN_X; x <= MAX_X; x++ {
		for y := MIN_Y; y <= MAX_Y; y++ {
			currentPower := calculatePowerOfArea(x, y, serialNumber)
			if currentPower > maxPower {
				maxPower = currentPower
				maxX = x
				maxY = y
			}
		}
	}
	fmt.Printf("Max power = %d at (%d, %d)\n", maxPower, maxX, maxY)
}
