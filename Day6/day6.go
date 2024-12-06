package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("AoC Day 6")
	// inputString := fileToString()

	testInput := `....#.....
	.........#
	..........
	..#.......
	.......#..
	..........
	.#..^.....
	........#.
	#.........
	......#...`

	// part 1
	lines := strings.Fields(testInput)

	grid := buildGrid(lines)
	printGrid(grid)

	start := getStartPos(grid)
	fmt.Println("Start pos:", start)

	pathGrid(grid, start)

	printGrid(grid)

	stepCount := 0
	for _, row := range grid {
		for _, col := range row {
			if col == "X" {
				stepCount++
			}
		}
	}
	fmt.Println("Num steps:", stepCount)
	fmt.Println(" ")

	// part 2
	linesPart2 := strings.Fields(testInput)
	newGrid := buildGrid(linesPart2)
	printGrid(newGrid)
	start2 := getStartPos(newGrid)
	fmt.Println("Start pos (still):", start2)

	// add obstacle
	newGrid[7][6] = "#"
	pathGrid(newGrid, start2)

	printGrid(newGrid)
}

func fileToString() string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	builder := new(strings.Builder)
	io.Copy(builder, file)
	return builder.String()
}

func buildGrid(lines []string) [][]string {
	size := len(lines)
	grid := make([][]string, size)
	for g := range grid {
		grid[g] = make([]string, size)
	}
	for i, line := range lines {
		chars := strings.Split(line, "")
		for j, char := range chars {
			grid[i][j] = char
		}
	}
	return grid
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func getStartPos(grid [][]string) []int {
	start := []int{0, 0}
	for i, row := range grid {
		for j, col := range row {
			if col == "^" {
				start = []int{i, j}
			}
		}
	}
	return start
}

func pathGrid(grid [][]string, start []int) {
	curr := start
	direction := "U"
	count := 0
	for direction != "END" {
		switch direction {
		case "U":
			direction = moveUp(grid, curr)
		case "D":
			direction = moveDown(grid, curr)
		case "R":
			direction = moveRight(grid, curr)
		case "L":
			direction = moveLeft(grid, curr)
		}
		count++
		if count > 5000 {
			fmt.Println("Inf")
			break
		}
	}
	fmt.Println("End pos:", curr, "End direction:", direction)
}

func moveUp(grid [][]string, curr []int) string {
	for grid[curr[0]][curr[1]] != "#" {
		grid[curr[0]][curr[1]] = "X"
		curr[0]--
		if curr[0] < 0 {
			return "END"
		}
	}
	curr[0]++
	return "R"
}

func moveRight(grid [][]string, curr []int) string {
	for grid[curr[0]][curr[1]] != "#" {
		grid[curr[0]][curr[1]] = "X"
		curr[1]++
		if curr[1] > len(grid)-1 {
			return "END"
		}
	}
	curr[1]--
	return "D"
}

func moveDown(grid [][]string, curr []int) string {
	for grid[curr[0]][curr[1]] != "#" {
		grid[curr[0]][curr[1]] = "X"
		curr[0]++
		if curr[0] > len(grid)-1 {
			return "END"
		}
	}
	curr[0]--
	return "L"
}

func moveLeft(grid [][]string, curr []int) string {
	for grid[curr[0]][curr[1]] != "#" {
		grid[curr[0]][curr[1]] = "X"
		curr[1]--
		if curr[1] < 0 {
			return "END"
		}
	}
	curr[1]++
	return "U"
}
