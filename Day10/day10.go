package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type visitPos struct {
	pos     [2]int
	visited bool
}

func main() {
	fmt.Println("AoC Day 10")
	inputString := fileToString()

	// 	testInput := `0123
	// 1234
	// 8765
	// 9876`

	// 	testInput := `89010123
	// 78121874
	// 87430965
	// 96549874
	// 45678903
	// 32019012
	// 01329801
	// 10456732`

	// part 1
	lines := strings.Fields(inputString)

	grid := buildGrid(lines)
	//printGrid(grid)
	allStarts := findAllInGrid(grid, "0")
	allEnds := findAllInGrid(grid, "9")
	//fmt.Println(allStarts)

	var visitMap []visitPos
	for _, end := range allEnds {
		vp := visitPos{pos: end, visited: false}
		visitMap = append(visitMap, vp)
	}

	totalCount := 0
	numSet := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, start := range allStarts {
		trailMapVisited := slices.Clone(visitMap)
		findNextStep(start, grid, numSet, 0, trailMapVisited)
		count := 0
		for _, visit := range trailMapVisited {
			if visit.visited {
				count++
			}
		}
		// fmt.Println("Start", start, "number of trails", count)
		totalCount += count
	}

	fmt.Println("Total", totalCount)

	// part 2

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
	fmt.Println(" ")
}

func findInGrid(grid [][]string, target string) [2]int {
	var res [2]int
	for i, row := range grid {
		for j, col := range row {
			if col == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}

func findAllInGrid(grid [][]string, target string) [][2]int {
	var res [][2]int
	for i, row := range grid {
		for j, col := range row {
			if col == target {
				var curr [2]int
				curr[0], curr[1] = i, j
				res = append(res, curr)
			}
		}
	}
	return res
}

func getInDirection(start [2]int, target string, grid [][]string, i int, j int) [2]int {
	grindLen := len(grid)
	if start[0]+i > -1 && start[0]+i < grindLen && start[1]+j > -1 && start[1]+j < grindLen && grid[start[0]+i][start[1]+j] == target {
		var curr [2]int
		curr[0], curr[1] = start[0]+i, start[1]+j
		return curr
	}
	return [2]int{-1, -1}
}

func findNextStep(start [2]int, grid [][]string, targets []string, targetIndex int, trailMapVisited []visitPos) {
	if targetIndex == 9 {
		for i := 0; i < len(trailMapVisited); i++ {
			if trailMapVisited[i].pos[0] == start[0] && trailMapVisited[i].pos[1] == start[1] {
				trailMapVisited[i].visited = true
			}
		}
		return
	}
	var nextSteps [][2]int
	up := getInDirection(start, targets[targetIndex], grid, -1, 0)
	if up[0] != -1 {
		// fmt.Println(targets[targetIndex], "at", up)
		nextSteps = append(nextSteps, up)
	}
	down := getInDirection(start, targets[targetIndex], grid, 1, 0)
	if down[0] != -1 {
		// fmt.Println(targets[targetIndex], "at", down)
		nextSteps = append(nextSteps, down)
	}
	left := getInDirection(start, targets[targetIndex], grid, 0, -1)
	if left[0] != -1 {
		// fmt.Println(targets[targetIndex], "at", left)
		nextSteps = append(nextSteps, left)
	}
	right := getInDirection(start, targets[targetIndex], grid, 0, 1)
	if right[0] != -1 {
		// fmt.Println(targets[targetIndex], "at", right)
		nextSteps = append(nextSteps, right)
	}

	if len(nextSteps) == 0 {
		return
	}

	targetIndex++
	for _, step := range nextSteps {
		findNextStep(step, grid, targets, targetIndex, trailMapVisited)
	}
}
