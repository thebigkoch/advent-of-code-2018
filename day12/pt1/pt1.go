package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getPositionElement(state string, position int) string {
	if position < 0 {
		return "."
	} else if position > len(state)-1 {
		return "."
	} else {
		return string(state[position])
	}
}

func getPositionString(state string, position int) string {
	positionString := ""
	for i := -2; i <= 2; i++ {
		positionString += getPositionElement(state, position+i)
	}
	return positionString
}

func processPosition(state string, rules map[string]string, position int) string {
	var result string
	var contains bool

	positionString := getPositionString(state, position)
	result, contains = rules[positionString]

	if !contains {
		result = "."
	}
	return result
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
	initialState := initialStateMessage[15:]
	initialState = "........................." + initialState + "........................."
	fmt.Printf(" 0: %s\n", initialState)

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

	var iterations int = 20
	startState := initialState
	for i := 1; i <= iterations; i++ {
		endState := ""
		for position := 0; position < len(startState); position++ {
			currentPos := processPosition(startState, rules, position)
			endState += currentPos
		}
		startState = endState
		fmt.Printf("%2d: %s\n", i, startState)
	}

	currentPos := -25
	total := 0
	for _, val := range startState {
		if val == '#' {
			total += currentPos
		}
		currentPos++
	}

	fmt.Printf("Total value: %d\n", total)
}
