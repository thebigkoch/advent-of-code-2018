package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/thebigkoch/advent-of-code-2018/day9/utils/circular_list"
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

	// Read input file.
	var fileContents string
	var numOfPlayers, numOfPoints int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fileContents = scanner.Text()
	}

	fmt.Sscanf(fileContents, "%d players; last marble is worth %d points", &numOfPlayers, &numOfPoints)
	numOfPoints *= 100

	// Initialize all player scores to 0.
	var playerScores []int
	playerScores = make([]int, numOfPlayers)
	for i := 0; i < numOfPlayers; i++ {
		playerScores[i] = 0
	}

	// Initialize the circle.
	var firstNode *circular_list.CircleNode = &circular_list.CircleNode{}
	firstNode.Value = 0
	firstNode.Prev = firstNode
	firstNode.Next = firstNode

	var currentNode *circular_list.CircleNode = firstNode
	var currentPlayer int = 0

	for i := 1; i <= numOfPoints; i++ {
		if i%23 == 0 {
			// Every 23rd turn, we add the current marble to the current player's score.
			playerScores[currentPlayer] += i

			// Additionally, we add the marble that's 7 left of the current marble.
			currentNode, _ = currentNode.GetPosition(-7)
			playerScores[currentPlayer] += currentNode.Value

			// Then we remove that marble.
			currentNode.Prev.Next = currentNode.Next
			currentNode = currentNode.Next
		} else {
			currentNode, _ = currentNode.GetPosition(1)
			currentNode.InsertNext(i)
			currentNode = currentNode.Next
		}
		currentPlayer++
		currentPlayer = currentPlayer % numOfPlayers
	}

	var highestScore int = 0
	for i := 0; i < numOfPlayers; i++ {
		fmt.Printf("Player #%d: %d pts\n", i, playerScores[i])
		if highestScore < playerScores[i] {
			highestScore = playerScores[i]
		}
	}
	fmt.Printf("Highest score: %d pts\n", highestScore)
}
