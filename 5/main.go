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

func reorderLine(updateLine []int, pageOrderingLine [][]int) []int {
	orderedPairs := make(map[int][]int)
	for _, pair := range pageOrderingLine {
		firstNum, secondNum := pair[0], pair[1]
		orderedPairs[firstNum] = append(orderedPairs[firstNum], secondNum)
	}

	for {
		swapped := false
		for i := 0; i < len(updateLine)-1; i++ {
			for _, secondNum := range orderedPairs[updateLine[i]] {
				if updateLine[i+1] == secondNum {
					updateLine[i], updateLine[i+1] = updateLine[i+1], updateLine[i]
					swapped = true
				}
			}
		}
		if !swapped {
			break
		}
	}

	return updateLine
}

func findMiddleUpdate(updateLine [][]int) int {
	sum := 0
	for i := 0; i < len(updateLine); i++ {
		n := len(updateLine[i])
		if n == 0 {
			break
		}
		midIndex := n / 2
		if n%2 == 0 {
			midIndex = midIndex - 1
			sum += updateLine[i][midIndex]
		}
		sum += updateLine[i][midIndex]
	}
	return sum
}

func parseUpdate(updateLine []int, firstNum int, secondNum int) bool {
	index1, index2 := -1, -1
	for i, v := range updateLine {
		if v == firstNum && index1 == -1 {
			index1 = i
		}
		if v == secondNum && index2 == -1 {
			index2 = i
		}
	}
	if index1 == -1 || index2 == -1 {
		return true
	}
	if index1 < index2 {
		return true
	}
	return false
}

func parseUpdates(pageOrderingInt [][]int, updatesInt [][]int) (int, int) {
	updatesLen := len(updatesInt)
	pageOrderingLen := len(pageOrderingInt)
	var correctLines [][]int
	var wrongLines [][]int
	correctMiddles := 0
	wrongMiddles := 0
	for i := 0; i < updatesLen; i++ {
		deadLine := false
		updateLine := updatesInt[i]
		for j := 0; j < pageOrderingLen; j++ {
			pageOrderingLine := pageOrderingInt[j]
			firstNum := pageOrderingLine[0]
			secondNum := pageOrderingLine[1]
			if !parseUpdate(updateLine, firstNum, secondNum) {
				deadLine = true
				updateLine = reorderLine(updateLine, pageOrderingInt)
				break
			}
		}
		if deadLine {
			wrongLines = append(wrongLines, updateLine)
		} else {
			correctLines = append(correctLines, updateLine)
		}
	}

	correctMiddles += findMiddleUpdate(correctLines)

	for _, wrongLine := range wrongLines {
		wrongMiddles += findMiddleUpdate([][]int{wrongLine})
	}

	return correctMiddles, wrongMiddles
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

	fmt.Println(parseUpdates(pageOrderingInt, updatesInt))
}
