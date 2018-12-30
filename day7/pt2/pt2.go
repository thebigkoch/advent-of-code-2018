package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
)

type node struct {
	predecessors []string
	followers    []string
}

type nodeInProgress struct {
	nodeName      string
	remainingWork int
}

const numberOfWorkers = 5
const baseTime = 60

var totalTime int = 0

func getTimeForRune(r rune) int {
	time := baseTime + int(r) - int('A') + 1
	return time
}

func findNodesWithoutPredecessors(nodes map[string]node) []nodeInProgress {
	var results []string

	// Go through all nodes.  Find any that don't have predecessors.
	for index, n := range nodes {
		if len(n.predecessors) == 0 {
			results = append(results, index)
		}
	}

	if len(results) > numberOfWorkers {
		results = results[:numberOfWorkers]
	}

	var resultsInProgress []nodeInProgress
	for _, nodeName := range results {
		var newNode nodeInProgress
		newNode.nodeName = nodeName
		newNode.remainingWork = getTimeForRune(rune(nodeName[0]))
		resultsInProgress = append(resultsInProgress, newNode)
	}

	return resultsInProgress
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsNIP(s []nodeInProgress, e string) bool {
	for _, a := range s {
		if a.nodeName == e {
			return true
		}
	}
	return false
}

func processNodesInProgress(nodesInProgress []nodeInProgress, nodesVisited []string) ([]nodeInProgress, []string) {
	var resultNodesInProgress []nodeInProgress
	var resultNodesVisited []string
	var newNodesVisited []string

	resultNodesVisited = nodesVisited

	lowestRemainingWork := math.MaxInt32
	for _, n := range nodesInProgress {
		if n.remainingWork < lowestRemainingWork {
			lowestRemainingWork = n.remainingWork
		}
	}

	totalTime = totalTime + lowestRemainingWork

	for _, n := range nodesInProgress {
		remainingWork := n.remainingWork - lowestRemainingWork
		if remainingWork == 0 {
			newNodesVisited = append(newNodesVisited, n.nodeName)
		} else {
			n.remainingWork = remainingWork
			resultNodesInProgress = append(resultNodesInProgress, n)
		}
	}

	sort.Strings(newNodesVisited)
	for _, s := range newNodesVisited {
		resultNodesVisited = append(resultNodesVisited, s)
	}

	return resultNodesInProgress, resultNodesVisited
}

func processRemainingNodes(nodesVisited []string, nodes map[string]node) []string {
	var nodesReadyToVisit []string

	for index, n := range nodes {
		// If this node was already visited, don't put it in the list of nodes ready to visit.
		if contains(nodesVisited, index) == false {
			allPredecessorsVisited := true

			// If any required predecessor was not already visited, then node n is not ready to visit.
			for _, predecessor := range n.predecessors {
				if contains(nodesVisited, predecessor) == false {
					allPredecessorsVisited = false
				}
			}

			// If all predecessors were visited, then node n is ready to visit.
			if allPredecessorsVisited {
				nodesReadyToVisit = append(nodesReadyToVisit, index)
			}
		}
	}

	return nodesReadyToVisit
}

func startNewNodes(nodesInProgress []nodeInProgress, nodesVisited []string, nodes map[string]node) []nodeInProgress {
	var resultNodesInProgress []nodeInProgress

	resultNodesInProgress = nodesInProgress
	nodesReadyToVisit := processRemainingNodes(nodesVisited, nodes)
	sort.Strings(nodesReadyToVisit)

	for numberOfWorkers > len(resultNodesInProgress) && len(nodesReadyToVisit) > 0 {
		var newNodeInProgress nodeInProgress
		nodeName := nodesReadyToVisit[0]
		nodesReadyToVisit = nodesReadyToVisit[1:]
		if containsNIP(resultNodesInProgress, nodeName) == false {
			newNodeInProgress.nodeName = nodeName
			newNodeInProgress.remainingWork = getTimeForRune(rune(nodeName[0]))
			resultNodesInProgress = append(resultNodesInProgress, newNodeInProgress)
		}
	}

	return resultNodesInProgress
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

	// Initialize our map of nodes
	nodes := make(map[string]node)

	// Read all points from the input file.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		message := scanner.Text()
		var predecessor, follower string
		_, err := fmt.Sscanf(message, "Step %1s must be finished before step %1s can begin.", &predecessor, &follower)
		if err != nil {
			log.Fatal(err)
		}

		p, _ := nodes[predecessor]
		p.followers = append(p.followers, follower)
		nodes[predecessor] = p

		f, _ := nodes[follower]
		f.predecessors = append(f.predecessors, predecessor)
		nodes[follower] = f
	}

	// Find all nodes without predecessors
	nodesInProgress := findNodesWithoutPredecessors(nodes)

	var nodesVisited []string
	for len(nodesVisited) < len(nodes) {
		nodesInProgress, nodesVisited = processNodesInProgress(nodesInProgress, nodesVisited)
		nodesInProgress = startNewNodes(nodesInProgress, nodesVisited, nodes)
	}

	fmt.Printf("Route: %v\n", nodesVisited)
	fmt.Printf("Total time: %d\n", totalTime)
}
