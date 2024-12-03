package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mul(a int, b int) int {
	return a * b
}

func sumMul(matches [][]string) int {
	mulSum := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])

		result := mul(a, b)

		fmt.Printf("mul(%d, %d) = %d\n", a, b, result)
		mulSum += result
	}
	return mulSum
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
	var result []string

	parts := strings.Split(inputString, "don't()")

	//Couldn't figure out a regex for this :(
	for i, part := range parts {
		if i == 0 {
			result = append(result, part)
		} else {
			if idx := strings.Index(part, "do()"); idx != -1 {
				result = append(result, part[idx+len("do()"):])
			}
		}
	}

	onlyDos := strings.Join(result, "")

	validRegex := `mul\(([0-9]{1,3}),([0-9]{1,3})\)`
	re := regexp.MustCompile(validRegex)

	// doRegex := regexp.MustCompile(`do\(\)([^d]+(?:d(?!on't\().)*)`)

	// splitString := doRegex.Split(inputString, -1)

	// fmt.Println(splitString)

	matches := re.FindAllStringSubmatch(onlyDos, -1)

	fmt.Println(sumMul(matches))
}
