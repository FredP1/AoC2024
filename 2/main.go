package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func safe_check_2(reports []int) int {
	absDiff := 0

	for i := 0; i < len(reports)-1; i++ {
		absDiff = AbsDiff(reports[i], reports[i+1])
		if absDiff > 3 {
			return 0
		}
		if absDiff == 0 {
			return 0
		}
	}
	return 1
}

func safe_check_1(reports []int) int {
	increaseFlag := 0
	decreaseFlag := 0
	for i := 0; i < len(reports)-1; i++ {
		if reports[i] > reports[i+1] {
			increaseFlag = 1
		} else if reports[i] < reports[i+1] {
			decreaseFlag = 1
		}
	}
	if increaseFlag == 1 && decreaseFlag == 1 {
		return 0
	}
	return 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var reports [][]int
	safe := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		var row []int
		for _, v := range values {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error converting value to int:", err)
				return
			}
			row = append(row, num)
		}
		reports = append(reports, row)
	}

	for i := 0; i < len(reports); i++ {
		check1 := safe_check_1(reports[i])
		check2 := safe_check_2(reports[i])

		if check1 == 1 && check2 == 1 {
			safe++
		}
	}
	fmt.Println(safe)
}
