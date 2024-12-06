package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertTo2DIntArray(strArray [][]string) ([][]int, error) {
	var intArray [][]int

	for _, row := range strArray {
		intRow := []int{}
		for _, str := range row {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("error converting '%s' to int: %v", str, err)
			}
			intRow = append(intRow, num)
		}
		intArray = append(intArray, intRow)
	}

	return intArray, nil
}

func createPageOrder(pageOrderingString [][]int) []int {
	var sortedPages []int
	for i := 0; i < len(pageOrderingString); i++ {

	}
	return sortedPages

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var pageOrdering [][]string
	var updates [][]string

	updatesFlag := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			updatesFlag = true
			continue
		}

		if updatesFlag {
			updates = append(updates, strings.Split(line, ","))
		} else {
			pageOrdering = append(pageOrdering, strings.Split(line, "|"))
		}
	}

	// Convert string arrays to int arrays
	pageOrderingInt, err := convertTo2DIntArray(pageOrdering)
	if err != nil {
		fmt.Println("Error converting pageOrdering to integers:", err)
		return
	}

	updatesInt, err := convertTo2DIntArray(updates)
	if err != nil {
		fmt.Println("Error converting updates to integers:", err)
		return
	}

	pagesOrdered := createPageOrder(pageOrderingInt)

	fmt.Println(pagesOrdered)
	fmt.Println(updatesInt)
}
