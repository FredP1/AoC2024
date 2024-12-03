package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func mul(a int, b int) int {
	return a * b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var inputString string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString += scanner.Text() + " "
	}

	mulSum := 0
	validRegex := `mul\(([0-9]{1,3}),([0-9]{1,3})\)`
	re := regexp.MustCompile(validRegex)

	matches := re.FindAllStringSubmatch(inputString, -1)

	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])

		result := mul(a, b)

		fmt.Printf("mul(%d, %d) = %d\n", a, b, result)
		mulSum += result
	}
	fmt.Println(mulSum)
}
