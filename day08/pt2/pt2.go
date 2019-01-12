package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type node struct {
	children []node
	metadata []int
}

// Returns:
//	node: The created node, with associated child nodes and metadata
//  string[]: Any remaining portion of the input string.
func convertNumbersToNode(numbersList []string) (node, []string) {
	var childrenCount int
	var metadataCount int
	childrenCount, _ = strconv.Atoi(numbersList[0])
	metadataCount, _ = strconv.Atoi(numbersList[1])
	var remainingNumbers []string = numbersList[2:]

	fmt.Printf("Children count: %d\n", childrenCount)
	fmt.Printf("Metadata count: %d\n\n", metadataCount)

	var currentNode node
	for i := 0; i < childrenCount; i++ {
		var child node
		child, remainingNumbers = convertNumbersToNode(remainingNumbers)
		currentNode.children = append(currentNode.children, child)
	}

	for j := 0; j < metadataCount; j++ {
		var metadataElement int
		metadataElement, _ = strconv.Atoi(remainingNumbers[0])
		remainingNumbers = remainingNumbers[1:]
		currentNode.metadata = append(currentNode.metadata, metadataElement)
	}

	return currentNode, remainingNumbers
}

func getMetadataTotal(currentNode node) int {

	metadataTotal := 0

	if len(currentNode.children) == 0 {
		for i := 0; i < len(currentNode.metadata); i++ {
			metadataTotal += currentNode.metadata[i]
		}
	} else {
		for i := 0; i < len(currentNode.metadata); i++ {
			nodeNumber := currentNode.metadata[i]
			if nodeNumber <= len(currentNode.children) {
				metadataTotal += getMetadataTotal(currentNode.children[nodeNumber-1])
			}
		}
	}

	return metadataTotal
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
	var fileContents string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fileContents = scanner.Text()
	}

	numbersList := strings.Split(fileContents, " ")
	root, remainingNumbers := convertNumbersToNode(numbersList)

	if len(remainingNumbers) > 0 {
		os.Exit(1)
	}

	metadataTotal := getMetadataTotal(root)

	fmt.Printf("Metadata total: %d\n", metadataTotal)
}
