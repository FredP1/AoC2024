package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertInputToMap(fileName string) map[int][]int {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	data := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), "\n")
		for _, line := range lines {
			parts := strings.Split(line, ": ")
			if len(parts) == 2 {
				var key int
				fmt.Sscanf(parts[0], "%d", &key)

				// Parse the values
				valueParts := strings.Fields(parts[1])
				var values []int
				for _, value := range valueParts {
					var num int
					fmt.Sscanf(value, "%d", &num)
					values = append(values, num)
				}
				data[key] = values
			}
		}
	}

	return data
}

func generatePermutations(values []int) []int {
	var results []int
	operators := []string{"+", "*"}
	n := len(values)

	combinations := generateOperatorCombinations(n-1, operators)

	for _, combination := range combinations {
		results = append(results, evaluateExpression(values, combination))
	}
	return results

}

func generateOperatorCombinations(n int, operators []string) [][]string {
	if n == 0 {
		return [][]string{{}}
	}
	smaller := generateOperatorCombinations(n-1, operators)
	var combinations [][]string

	for _, combo := range smaller {
		for _, op := range operators {
			newCombo := append([]string{}, combo...)
			newCombo = append(newCombo, op)
			combinations = append(combinations, newCombo)
		}
	}
	return combinations
}

func evaluateExpression(arr []int, operators []string) int {
	result := arr[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			result += arr[i+1]
		case "*":
			result *= arr[i+1]
		}
	}
	return result
}

func part1Test(inputMap map[int][]int) int {

	var correctValues int
	for key, values := range inputMap {
		matched := false
		results := generatePermutations(values)
		for _, result := range results {
			if result == key {
				matched = true
			}
		}

		if matched {
			fmt.Printf("Target %d can be achieved with the array %v\n", key, values)
			correctValues += key
		} else {
			fmt.Printf("Target %d cannot be achieved with the array %v\n", key, values)
		}
	}
	return correctValues
}

func main() {
	fileName := "input.txt"
	inputMap := convertInputToMap(fileName)
	fmt.Println(part1Test(inputMap))

}
