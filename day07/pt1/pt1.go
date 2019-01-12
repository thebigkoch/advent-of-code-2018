package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type node struct {
	predecessors []string
	followers    []string
}

func findNodesWithoutPredecessors(nodes map[string]node) []string {
	var results []string

	// Go through all nodes.  Find any that don't have predecessors.
	for index, n := range nodes {
		if len(n.predecessors) == 0 {
			results = append(results, index)
		}
	}

	return results
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
	nodesReadyToVisit := findNodesWithoutPredecessors(nodes)

	var nodesVisited []string
	for len(nodesReadyToVisit) > 0 {
		sort.Strings(nodesReadyToVisit)
		nodesVisited = append(nodesVisited, nodesReadyToVisit[0])
		nodesReadyToVisit = processRemainingNodes(nodesVisited, nodes)
	}

	fmt.Printf("Route: %v", nodesVisited)
}
