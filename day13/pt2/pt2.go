package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type cartType struct {
	x         int
	y         int
	direction rune
	nextTurn  int
	collided  bool
}

func getNewDirection(direction rune, nextTurn int) rune {
	var returnVal rune
	if direction == '<' {
		if nextTurn%3 == 0 {
			returnVal = 'v'
		} else if nextTurn%3 == 1 {
			returnVal = direction
		} else if nextTurn%3 == 2 {
			returnVal = '^'
		}
	} else if direction == '>' {
		if nextTurn%3 == 0 {
			returnVal = '^'
		} else if nextTurn%3 == 1 {
			returnVal = direction
		} else if nextTurn%3 == 2 {
			returnVal = 'v'
		}
	} else if direction == '^' {
		if nextTurn%3 == 0 {
			returnVal = '<'
		} else if nextTurn%3 == 1 {
			returnVal = direction
		} else if nextTurn%3 == 2 {
			returnVal = '>'
		}
	} else if direction == 'v' {
		if nextTurn%3 == 0 {
			returnVal = '>'
		} else if nextTurn%3 == 1 {
			returnVal = direction
		} else if nextTurn%3 == 2 {
			returnVal = '<'
		}
	} else {
		log.Fatal("Error! Invalid direction.")
		os.Exit(1)
	}

	return returnVal
}

func sortCarts(carts []cartType) []cartType {
	for i := 0; i < len(carts); i++ {
		for j := i + 1; j < len(carts); j++ {
			if (carts[i].y > carts[j].y) || (carts[i].y == carts[j].y && carts[i].x > carts[j].x) {
				carts[i], carts[j] = carts[j], carts[i]
			}
		}
	}

	return carts
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

	// Read input.
	grid := make(map[int]string)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid[i] = line
		i++
	}

	// Find the places on the track where we have carts.  Track these locations.
	var carts []cartType
	for idy, line := range grid {
		lineAsRune := []rune(line)
		for idx, r := range line {
			var cart cartType
			if r == '<' || r == '>' || r == '^' || r == 'v' {
				cart.x = idx
				cart.y = idy
				cart.direction = r
				cart.nextTurn = 0
				cart.collided = false
				carts = append(carts, cart)

				if r == '<' || r == '>' {
					lineAsRune[idx] = '-'
				} else {
					lineAsRune[idx] = '|'
				}
			}
		}
		grid[idy] = string(lineAsRune)
	}

	remainingCarts := len(carts)
	for {
		carts = sortCarts(carts)
		for idx, cart := range carts {
			if cart.collided == false {
				// Update the cart location, based on its current direction.
				dir := cart.direction
				if dir == '<' {
					cart.x = cart.x - 1
				} else if dir == '>' {
					cart.x = cart.x + 1
				} else if dir == '^' {
					cart.y = cart.y - 1
				} else if dir == 'v' {
					cart.y = cart.y + 1
				}

				// Check for collision
				for idy, cart2 := range carts {
					if idx != idy && cart2.collided == false {
						if cart.x == cart2.x && cart.y == cart2.y {
							fmt.Printf("Collision between cart #%d and cart #%d\n", idx, idy)
							fmt.Printf("Collision at point (%d, %d)\n", cart.x, cart.y)
							cart.collided = true
							cart2.collided = true
							carts[idy] = cart2
							remainingCarts = remainingCarts - 2
						}
					}
				}

				// Check if we need to turn the cart.
				if cart.collided == false {
					currentTrack := grid[cart.y][cart.x]
					if currentTrack == '/' {
						if cart.direction == '>' {
							cart.direction = '^'
						} else if cart.direction == 'v' {
							cart.direction = '<'
						} else if cart.direction == '<' {
							cart.direction = 'v'
						} else if cart.direction == '^' {
							cart.direction = '>'
						} else {
							log.Fatal("Error: Invalid drection.")
							os.Exit(1)
						}
					} else if currentTrack == '\\' {
						if cart.direction == '>' {
							cart.direction = 'v'
						} else if cart.direction == 'v' {
							cart.direction = '>'
						} else if cart.direction == '<' {
							cart.direction = '^'
						} else if cart.direction == '^' {
							cart.direction = '<'
						} else {
							log.Fatal("Error: Invalid drection.")
							os.Exit(1)
						}
					} else if currentTrack == '+' {
						cart.direction = getNewDirection(cart.direction, cart.nextTurn)
						cart.nextTurn++
					}
				}

				carts[idx] = cart
			}
		}

		if remainingCarts == 1 {
			for _, cart := range carts {
				if cart.collided == false {
					fmt.Printf("Last cart at (%d,%d)\n", cart.x, cart.y)
					return
				}
			}
		}
	}
}
