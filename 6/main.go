package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Guard struct {
	movementDirection int
	xPosition         int
	yPosition         int
}

func (g *Guard) checkForObstacle(grid [][]rune) bool {
	hashRune := '#'
	if g.movementDirection == 1 { //North
		if grid[g.yPosition-1][g.xPosition] == hashRune {
			return true
		}
	} else if g.movementDirection == 2 { //East
		if grid[g.yPosition][g.xPosition+1] == hashRune {
			return true
		}
	} else if g.movementDirection == 3 { //South
		if grid[g.yPosition+1][g.xPosition] == hashRune {
			return true
		}
	} else if g.movementDirection == 4 { //West
		if grid[g.yPosition][g.xPosition-1] == hashRune {
			return true
		}
	}
	return false
}

func (g *Guard) boundsCheck(grid [][]rune) bool {
	if g.movementDirection == 1 && g.yPosition-1 < 0 { //North
		return true
	}
	if g.movementDirection == 2 && g.xPosition+2 > len(grid) { //East
		return true
	}
	if g.movementDirection == 3 && g.yPosition+2 > len(grid[0]) { //South
		return true
	}
	if g.movementDirection == 4 && g.xPosition-1 < 0 { //West
		return true
	}
	return false
}

func (g *Guard) move(grid [][]rune) [][]rune {
	if g.movementDirection == 1 { //North
		if !g.checkForObstacle(grid) {
			grid[g.yPosition][g.xPosition] = 'X'
			g.yPosition--
			grid[g.yPosition][g.xPosition] = '^'
		} else {
			g.movementDirection = 2
		}
	} else if g.movementDirection == 2 { //East
		if !g.checkForObstacle(grid) {
			grid[g.yPosition][g.xPosition] = 'X'
			g.xPosition++
			grid[g.yPosition][g.xPosition] = '^'
		} else {
			g.movementDirection = 3
		}
	} else if g.movementDirection == 3 { //South
		if !g.checkForObstacle(grid) {
			grid[g.yPosition][g.xPosition] = 'X'
			g.yPosition++
			grid[g.yPosition][g.xPosition] = '^'
		} else {
			g.movementDirection = 4
		}
	} else if g.movementDirection == 4 { //West
		if !g.checkForObstacle(grid) {
			grid[g.yPosition][g.xPosition] = 'X'
			g.xPosition--
			grid[g.yPosition][g.xPosition] = '^'
		} else {
			g.movementDirection = 1
		}
	}
	return grid
}

func findGuardPosition(grid [][]rune) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func convertTo2DArrayFromFile(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}

	return grid
}

func part1(grid [][]rune) int {
	var positionMap = make(map[Guard]int)
	guardYPosition, guardXPosition := findGuardPosition(grid)
	guard := Guard{movementDirection: 1, xPosition: guardXPosition, yPosition: guardYPosition}
	for 0 < len(grid) && 0 < len(grid[0]) {
		// fmt.Println("New Grid")
		if guard.boundsCheck(grid) {
			break
		}
		grid = guard.move(grid)
		positionMap[guard]++
		for _, value := range positionMap {
			if value > 1 {
				return 1
			}
		}
	}
	return 0
	// for _, row := range grid {
	// 	fmt.Println(string(row))
	// }
}

func deepCopyGrid(grid [][]rune) [][]rune {
	copiedGrid := make([][]rune, len(grid))
	for i := range grid {
		copiedGrid[i] = make([]rune, len(grid[i]))
		copy(copiedGrid[i], grid[i])
	}
	return copiedGrid
}

func part2(grid [][]rune) {
	hashRune := '#'
	guardRune := '^'
	gridOriginal := deepCopyGrid(grid)
	counter := 0
	var wg sync.WaitGroup

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != hashRune && grid[i][j] != guardRune {
				gridCopy := deepCopyGrid(gridOriginal) // Create copy per goroutine to avoid race

				wg.Add(1)
				go func(i, j int) {
					defer wg.Done()

					// Modify the grid in the goroutine
					gridCopy[i][j] = hashRune
					fmt.Println("Processing Row:", i, "Column:", j)
					counter += part1(gridCopy) // Update the counter safely

				}(i, j)
			}
		}
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println(counter)
}

func main() {
	fileName := "input.txt"

	grid := convertTo2DArrayFromFile(fileName)
	// for _, row := range grid {
	// 	fmt.Println(string(row))
	// }
	// part1(grid)
	part2(grid)
}
