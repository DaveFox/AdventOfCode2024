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
	fmt.Println("AoC Day 8")
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
	// 	testInput2 := `T.........
	// ...T......
	// .T........
	// ..........
	// ..........
	// ..........
	// ..........
	// ..........
	// ..........
	// ..........`

	// part 1
	lines := strings.Fields(inputString)

	grid := buildGrid(lines)
	printGrid(grid)

	antennas := getAntennas(grid)

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

	nodes = removeDuplicateNodes(nodes)

	addNodesToGrid(nodes, grid)
	printGrid(grid)
	fmt.Println("Antinode count:", len(nodes))

	// part 2
	grid2 := buildGrid(lines)
	printGrid(grid2)

	// find distance to other antenna of same type
	var nodes2 []antennaInfo
	for o, antOuter := range antennas {
		for i, antInner := range antennas {
			if antInner.freq == antOuter.freq && o != i {
				dx := antOuter.pos[0] - antInner.pos[0]
				dy := antOuter.pos[1] - antInner.pos[1]

				posX := antOuter.pos[0]
				posY := antOuter.pos[1]
				for posX < len(lines) && posX > -1 && posY < len(lines) && posY > -1 {
					nodeX := posX + dx
					nodeY := posY + dy
					if nodeX > -1 && nodeX < len(lines) && nodeY > -1 && nodeY < len(lines) {
						newNode := antennaInfo{freq: "#", pos: make([]int, 2)}
						newNode.pos[0] = nodeX
						newNode.pos[1] = nodeY
						nodes2 = append(nodes2, newNode)
					}
					posX = posX + dx
					posY = posY + dy
				}
			}
		}
	}

	nodes2 = removeDuplicateNodes(nodes2)

	addNodesToGrid(nodes2, grid2)
	printGrid(grid2)

	total := 0
	for _, row := range grid2 {
		for _, col := range row {
			if col != "." {
				total++
			}
		}
	}
	fmt.Println("Antinode count:", len(nodes2))
	fmt.Println("Grid dispaly count:", total)
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
		if grid[node.pos[0]][node.pos[1]] == "." {
			grid[node.pos[0]][node.pos[1]] = node.freq
		}
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
