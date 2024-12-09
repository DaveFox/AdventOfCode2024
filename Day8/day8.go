package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type antennaInfo struct {
	freq string
	pos  []int
}

func main() {
	fmt.Println("AoC Day 6")
	inputString := fileToString()

	// 	testInput := `............
	// ........0...
	// .....0......
	// .......0....
	// ....0.......
	// ......A.....
	// ............
	// ............
	// ........A...
	// .........A..
	// ............
	// ............`

	// part 1
	lines := strings.Fields(inputString)

	grid := buildGrid(lines)
	printGrid(grid)

	antennas := getAntennas(grid)
	// fmt.Println("Antenna list:", antennas)

	// find distance to other antenna of same type
	var nodes []antennaInfo
	for o, antOuter := range antennas {
		for i, antInner := range antennas {
			if antInner.freq == antOuter.freq && o != i {
				dx := antOuter.pos[0] - antInner.pos[0]
				dy := antOuter.pos[1] - antInner.pos[1]

				nodeX := antOuter.pos[0] + dx
				nodeY := antOuter.pos[1] + dy
				if nodeX >= 0 && nodeX < len(lines) && nodeY >= 0 && nodeY < len(lines) {
					newNode := antennaInfo{freq: "#", pos: make([]int, 2)}
					newNode.pos[0] = nodeX
					newNode.pos[1] = nodeY
					nodes = append(nodes, newNode)
				}
			}
		}
	}
	// fmt.Println(nodes)

	nodes = removeDuplicateNodes(nodes)
	// fmt.Println(nodes)
	// remove overlaps
	// var finalNodes []antennaInfo
	// for _, node := range nodes {
	// 	conflict := false
	// 	for _, antenna := range antennas {
	// 		if node.pos[0] == antenna.pos[0] && node.pos[1] == antenna.pos[1] {
	// 			conflict = true
	// 			break
	// 		}
	// 	}
	// 	if !conflict {
	// 		finalNodes = append(finalNodes, node)
	// 	}
	// }

	addNodesToGrid(nodes, grid)
	printGrid(grid)
	fmt.Println("Antinode count:", len(nodes))

	// part 2
	// printGrid(newGrid)
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

func getAntennas(grid [][]string) []antennaInfo {
	var antennas []antennaInfo
	for i, row := range grid {
		for j, col := range row {
			if col != "." {
				a := antennaInfo{freq: col, pos: make([]int, 2)}
				a.pos[0] = i
				a.pos[1] = j
				antennas = append(antennas, a)
			}
		}
	}
	return antennas
}

func addNodesToGrid(nodes []antennaInfo, grid [][]string) {
	for _, node := range nodes {
		grid[node.pos[0]][node.pos[1]] = node.freq
	}
}

func removeDuplicateNodes(nodes []antennaInfo) []antennaInfo {
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].pos[0] == nodes[j].pos[0] && nodes[i].pos[1] == nodes[j].pos[1] {
				nodes = slices.Delete(nodes, j, j+1)
				j--
			}
		}
	}
	return nodes
}
