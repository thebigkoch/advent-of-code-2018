package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/thebigkoch/advent-of-code-2018/utils/bst"
)

func main() {
	node := &bst.BinaryTreeNode{}

	// Get the current directory.
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Open the input file.
	filePath := filepath.Join(currentDir, "input.txt")
	var total int64 = 0
	var counter int = 0

	for {
		f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			i, err := strconv.ParseInt(scanner.Text(), 10, 32)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			total += i
			strTotal := fmt.Sprintf("%d", total)
			_, found := node.Find(strTotal)
			if found {
				fmt.Printf("Total: %d\n", total)
				os.Exit(0)
			}
			node.Insert(strTotal, strTotal)
		}
		counter++
		fmt.Printf("Finished loop #%d\n", counter)
	}
}
