package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func neCheck2(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx < 2 || (lineIdx > 137) {
		return 0
	} else if array[arrayIdx][lineIdx] != "M" {
		return 0
	} else if array[arrayIdx-1][lineIdx+1] != "A" {
		return 0
	} else if array[arrayIdx-2][lineIdx+2] != "S" {
		return 0
	}
	return 1
}

func nwCheck2(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx < 2 || lineIdx < 2 {
		return 0
	} else if array[arrayIdx][lineIdx] != "M" {
		return 0
	} else if array[arrayIdx-1][lineIdx-1] != "A" {
		return 0
	} else if array[arrayIdx-2][lineIdx-2] != "S" {
		return 0
	}
	return 1
}

func seCheck2(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx > 137 || lineIdx > 137 {
		return 0
	} else if array[arrayIdx][lineIdx] != "M" {
		return 0
	} else if array[arrayIdx+1][lineIdx+1] != "A" {
		return 0
	} else if array[arrayIdx+2][lineIdx+2] != "S" {
		return 0
	}
	return 1
}

func swCheck2(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx > 137 || lineIdx < 2 {
		return 0
	} else if array[arrayIdx][lineIdx] != "M" {
		return 0
	} else if array[arrayIdx+1][lineIdx-1] != "A" {
		return 0
	} else if array[arrayIdx+2][lineIdx-2] != "S" {
		return 0
	}
	return 1
}

func nCheck(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx < 2 {
		return 0
	} else if array[arrayIdx][lineIdx] != "X" {
		return 0
	} else if array[arrayIdx-1][lineIdx] != "M" {
		return 0
	} else if array[arrayIdx-2][lineIdx] != "A" {
		return 0
	} else if array[arrayIdx-3][lineIdx] != "S" {
		return 0
	}
	return 1
}

func sCheck(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx > 136 {
		return 0
	} else if array[arrayIdx][lineIdx] != "X" {
		return 0
	} else if array[arrayIdx+1][lineIdx] != "M" {
		return 0
	} else if array[arrayIdx+2][lineIdx] != "A" {
		return 0
	} else if array[arrayIdx+3][lineIdx] != "S" {
		return 0
	}
	return 1
}

func neCheck(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx < 2 || (lineIdx > 136) {
		return 0
	} else if array[arrayIdx][lineIdx] != "X" {
		return 0
	} else if array[arrayIdx-1][lineIdx+1] != "M" {
		return 0
	} else if array[arrayIdx-2][lineIdx+2] != "A" {
		return 0
	} else if array[arrayIdx-3][lineIdx+3] != "S" {
		return 0
	}
	return 1
}

func nwCheck(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx < 2 || lineIdx < 3 {
		return 0
	} else if array[arrayIdx][lineIdx] != "X" {
		return 0
	} else if array[arrayIdx-1][lineIdx-1] != "M" {
		return 0
	} else if array[arrayIdx-2][lineIdx-2] != "A" {
		return 0
	} else if array[arrayIdx-3][lineIdx-3] != "S" {
		return 0
	}
	return 1
}

func seCheck(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx > 136 || lineIdx > 136 {
		return 0
	} else if array[arrayIdx][lineIdx] != "X" {
		return 0
	} else if array[arrayIdx+1][lineIdx+1] != "M" {
		return 0
	} else if array[arrayIdx+2][lineIdx+2] != "A" {
		return 0
	} else if array[arrayIdx+3][lineIdx+3] != "S" {
		return 0
	}
	return 1
}

func swCheck(array [][]string, arrayIdx int, lineIdx int) int {
	if arrayIdx > 136 || lineIdx < 3 {
		return 0
	} else if array[arrayIdx][lineIdx] != "X" {
		return 0
	} else if array[arrayIdx+1][lineIdx-1] != "M" {
		return 0
	} else if array[arrayIdx+2][lineIdx-2] != "A" {
		return 0
	} else if array[arrayIdx+3][lineIdx-3] != "S" {
		return 0
	}
	return 1
}

func eCheck(array [][]string, arrayIdx int, lineIdx int) int {
	if lineIdx > 136 {
		return 0
	} else if array[arrayIdx][lineIdx] != "X" {
		return 0
	} else if array[arrayIdx][lineIdx+1] != "M" {
		return 0
	} else if array[arrayIdx][lineIdx+2] != "A" {
		return 0
	} else if array[arrayIdx][lineIdx+3] != "S" {
		return 0
	}
	return 1
}

func wCheck(array [][]string, arrayIdx int, lineIdx int) int {
	if lineIdx < 3 {
		return 0
	} else if array[arrayIdx][lineIdx] != "X" {
		return 0
	} else if array[arrayIdx][lineIdx-1] != "M" {
		return 0
	} else if array[arrayIdx][lineIdx-2] != "A" {
		return 0
	} else if array[arrayIdx][lineIdx-3] != "S" {
		return 0
	}
	return 1
}

func xmasCheck(array [][]string) int {
	total := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			total += nCheck(array, i, j)
			total += sCheck(array, i, j)
			total += neCheck(array, i, j)
			total += nwCheck(array, i, j)
			total += seCheck(array, i, j)
			total += swCheck(array, i, j)
			total += eCheck(array, i, j)
			total += wCheck(array, i, j)
		}
	}
	return total
}

func part2(array [][]string) int {
	total := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if neCheck2(array, i, j) == 1 {
				if seCheck2(array, i-2, j) == 1 {
					fmt.Println(array[i][j], array[i-1][j+1], array[i-2][j+2])
					fmt.Println(array[i-2][j], array[i-1][j+1], array[i-2][j+2])
					total++
				}
			}
			if nwCheck2(array, i, j) == 1 {
				if swCheck2(array, i-2, j) == 1 {
					total++
				}
			}

			if neCheck2(array, i, j) == 1 {
				if nwCheck2(array, i, j+2) == 1 {
					total++
				}
			}
			if seCheck2(array, i, j) == 1 {
				if swCheck2(array, i, j+2) == 1 {
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

	fmt.Println(xmasCheck(array))

	fmt.Println(part2(array))
}
