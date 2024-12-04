package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	RowStep, ColStep int
	Pattern          []string
}

func checkPattern(array [][]string, startRow, startCol int, direction Direction) bool {
	for step, char := range direction.Pattern {
		row := startRow + step*direction.RowStep
		col := startCol + step*direction.ColStep

		if row < 0 || row >= len(array) || col < 0 || col >= len(array[row]) {
			return false
		}

		if array[row][col] != char {
			return false
		}
	}
	return true
}

func xmasCheck(array [][]string) int {
	directions := []Direction{
		{RowStep: -1, ColStep: 0, Pattern: []string{"X", "M", "A", "S"}},  // North
		{RowStep: 1, ColStep: 0, Pattern: []string{"X", "M", "A", "S"}},   // South
		{RowStep: 0, ColStep: 1, Pattern: []string{"X", "M", "A", "S"}},   // East
		{RowStep: 0, ColStep: -1, Pattern: []string{"X", "M", "A", "S"}},  // West
		{RowStep: -1, ColStep: 1, Pattern: []string{"X", "M", "A", "S"}},  // Northeast
		{RowStep: -1, ColStep: -1, Pattern: []string{"X", "M", "A", "S"}}, // Northwest
		{RowStep: 1, ColStep: 1, Pattern: []string{"X", "M", "A", "S"}},   // Southeast
		{RowStep: 1, ColStep: -1, Pattern: []string{"X", "M", "A", "S"}},  // Southwest
	}
	total := 0

	for row := 0; row < len(array); row++ {
		for col := 0; col < len(array[row]); col++ {
			for _, dir := range directions {
				if checkPattern(array, row, col, dir) {
					total++
				}
			}
		}
	}

	return total
}

func part2(array [][]string) int {
	directions := map[string]Direction{
		"NE": {RowStep: -1, ColStep: 1, Pattern: []string{"M", "A", "S"}},
		"SE": {RowStep: 1, ColStep: 1, Pattern: []string{"M", "A", "S"}},
		"NW": {RowStep: -1, ColStep: -1, Pattern: []string{"M", "A", "S"}},
		"SW": {RowStep: 1, ColStep: -1, Pattern: []string{"M", "A", "S"}},
	}
	total := 0

	for row := 0; row < len(array); row++ {
		for col := 0; col < len(array[row]); col++ {
			if checkPattern(array, row, col, directions["NE"]) {
				if checkPattern(array, row-2, col, directions["SE"]) {
					total++
				}
			}
			if checkPattern(array, row, col, directions["NW"]) {
				if checkPattern(array, row-2, col, directions["SW"]) {
					total++
				}
			}
			if checkPattern(array, row, col, directions["NE"]) {
				if checkPattern(array, row, col+2, directions["NW"]) {
					total++
				}
			}
			if checkPattern(array, row, col, directions["SE"]) {
				if checkPattern(array, row, col+2, directions["SW"]) {
					total++
				}
			}
		}
	}

	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var array [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "")
		array = append(array, splitLine)
	}

	fmt.Println("Part 1:", xmasCheck(array))
	fmt.Println("Part 2:", part2(array))
}
