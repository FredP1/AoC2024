package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Antenna struct {
	character            string
	xLocation, yLocation int
}

type AntiNode struct {
	xLocation, yLocation int
}

func convertTo2DArrayFromFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	var grid [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		newLine := strings.Split(line, "")
		grid = append(grid, newLine)
	}

	return grid
}

func createAntinodes(pairs []Antenna, grid [][]string) {
	antinodeCounter := 0
	xDiff := pairs[0].xLocation - pairs[1].xLocation
	yDiff := pairs[0].yLocation - pairs[1].yLocation

	antiNodeX1 := pairs[0].xLocation + xDiff
	antiNodeY1 := pairs[0].yLocation + yDiff
	antiNode1 := AntiNode{xLocation: antiNodeX1, yLocation: antiNodeY1}

	antiNodeX2 := pairs[1].xLocation - xDiff
	antiNodeY2 := pairs[1].yLocation - yDiff
	antiNode2 := AntiNode{xLocation: antiNodeX2, yLocation: antiNodeY2}

	gridYLen := len(grid)
	gridXLen := len(grid[0])
	for (antiNode1.xLocation < gridXLen && antiNode1.yLocation < gridYLen) && (antiNode1.xLocation >= 0 && antiNode1.yLocation >= 0) {
		grid[antiNode1.yLocation][antiNode1.xLocation] = "#"
		antiNode1.xLocation = antiNode1.xLocation + xDiff
		antiNode1.yLocation = antiNode1.yLocation + yDiff
	}

	for (antiNode2.xLocation < gridXLen && antiNode2.yLocation < gridYLen) && (antiNode2.xLocation >= 0 && antiNode2.yLocation >= 0) {
		grid[antiNode2.yLocation][antiNode2.xLocation] = "#"
		antiNode2.xLocation = antiNode2.xLocation - xDiff
		antiNode2.yLocation = antiNode2.yLocation - yDiff
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != "." {
				antinodeCounter++
			}
		}
	}
	fmt.Println(antinodeCounter)
}

func pairAntennas(antennaArray []Antenna) [][]Antenna {
	var returnArray [][]Antenna
	for i := 0; i < len(antennaArray); i++ {
		for j := i + 1; j < len(antennaArray); j++ {
			pair := []Antenna{antennaArray[i], antennaArray[j]}
			returnArray = append(returnArray, pair)
		}
	}
	return returnArray
}

func groupAntennas(antennaArray []Antenna) map[string][]Antenna {
	antennaMap := make(map[string][]Antenna)
	for _, antenna := range antennaArray {
		antennaMap[antenna.character] = append(antennaMap[antenna.character], antenna)
	}
	return antennaMap
}

func findAntennas(grid [][]string) []Antenna {
	var antennas []Antenna
	for i, row := range grid {
		for j, column := range row {
			if column != "." {
				antenna := Antenna{character: column, xLocation: j, yLocation: i}
				antennas = append(antennas, antenna)
			}
		}
	}
	return antennas
}

func main() {
	fileName := "input.txt"

	grid := convertTo2DArrayFromFile(fileName)
	var antennaPairs [][]Antenna

	antennaMap := (findAntennas(grid))
	groupedAntennas := (groupAntennas(antennaMap))
	for _, values := range groupedAntennas {
		antennaPairs = append(antennaPairs, pairAntennas(values)...)
	}
	pairedAntennas := antennaPairs

	for _, values := range pairedAntennas {
		createAntinodes(values, grid)
	}

}
