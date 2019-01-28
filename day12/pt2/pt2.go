package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getPositionString(state map[int]bool, position int) string {
	positionString := ""

	var value bool = false
	var contains bool = false
	for i := position - 2; i <= position+2; i++ {
		value, contains = state[i]
		if !contains {
			value = false
		}
		if value {
			positionString += "#"
		} else {
			positionString += "."
		}
	}

	return positionString
}

func processPosition(state map[int]bool, rules map[string]string, position int) bool {
	var result string
	var contains bool

	positionString := getPositionString(state, position)
	result, contains = rules[positionString]

	if !contains {
		result = "."
	}

	if result == "." {
		return false
	} else {
		return true
	}
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
	scanner := bufio.NewScanner(f)

	// Read the first line.
	scanner.Scan()
	initialStateMessage := scanner.Text()
	initialStateMessage = initialStateMessage[15:]
	initialState := make(map[int]bool)
	leftPosition := 0
	rightPosition := 0
	for i := 0; i < len(initialStateMessage); i++ {
		if initialStateMessage[i] == '.' {
			initialState[i] = false
		} else if initialStateMessage[i] == '#' {
			initialState[i] = true
			rightPosition = i
		} else {
			fmt.Printf("Error!")
		}
	}

	// Read the second blank line.
	scanner.Scan()

	// Read all conditions / results.
	rules := make(map[string]string)
	for scanner.Scan() {
		message := scanner.Text()
		condition := message[0:5]
		result := string(message[9])
		rules[condition] = result
	}

	startState := initialState
	prevValue := ""
	iteration := 0
	for {
		endState := make(map[int]bool)
		bottom := leftPosition - 2
		top := rightPosition + 2
		for position := bottom; position <= top; position++ {
			currentPositionValue := processPosition(startState, rules, position)
			if position < leftPosition {
				if currentPositionValue {
					endState[position] = currentPositionValue
					leftPosition = position
				}
			} else if position > rightPosition {
				if currentPositionValue {
					endState[position] = currentPositionValue
					rightPosition = position
				}
			} else {
				endState[position] = currentPositionValue
			}
		}

		currentValue := ""
		for i := leftPosition; i <= rightPosition; i++ {
			value, contains := startState[i]
			if contains && value {
				currentValue += "#"
			} else {
				if currentValue != "" {
					currentValue += "."
				}
			}
		}

		startState = endState
		iteration++
		if currentValue == prevValue {
			break
		}
		prevValue = currentValue
	}

	adjustment := 50000000000 - iteration
	total := 0
	for position, value := range startState {
		if value == true {
			total = total + position + adjustment
		}
	}

	fmt.Printf("Total: %d", total)
}
