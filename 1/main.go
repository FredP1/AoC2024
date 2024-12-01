package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func second(leftColumn []int, rightColumn []int) {
	listMap := make(map[int]int)

	for i := 0; i < len(leftColumn); i++ {
		key := leftColumn[i]
		listMap[key] = 0
	}

	for i := 0; i < len(leftColumn); i++ {
		count := 0
		for _, v := range rightColumn {
			if v == leftColumn[i] {
				count++
			}
		}
		listMap[leftColumn[i]] = count
	}
	counter := 0
	for key, value := range listMap {
		result := key * value
		counter += result
	}
	fmt.Println(counter)

}

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var leftColumn []int
	var rightColumn []int
	var distance int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		columns := strings.Split(line, "   ")
		if len(columns) == 2 {
			leftValue, err1 := strconv.Atoi(strings.TrimSpace(columns[0]))
			rightValue, err2 := strconv.Atoi(strings.TrimSpace(columns[1]))
			if err1 != nil || err2 != nil {
				fmt.Println("Error converting line to integers:", err1, err2)
				continue
			}

			leftColumn = append(leftColumn, leftValue)
			rightColumn = append(rightColumn, rightValue)
		}
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for i := 0; i < len(leftColumn); i++ {
		localDistance := leftColumn[i] - rightColumn[i]
		if localDistance < 0 {
			localDistance = -localDistance
		}
		fmt.Printf("Distance between %d and %d: %d\n", leftColumn[i], rightColumn[i], localDistance)
		distance += int(localDistance)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(distance)
	second(leftColumn, rightColumn)
}
