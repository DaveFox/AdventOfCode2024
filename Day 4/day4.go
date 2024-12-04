package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("AoC Day 4")
	inputString := fileToString()

	// 	testIn := `MMMSXXMASM
	// MSAMXMSMSA
	// AMXSXMAAMM
	// MSAMASMSMX
	// XMASAMXAMM
	// XXAMMXXAMA
	// SMSMSASXSS
	// SAXAMASAAA
	// MAMMMXMMMM
	// MXMXAXMASX`

	// part 1
	rows := strings.Split(inputString, "\r\n") // \r\n or \n

	rowTotal := checkRows(rows)
	colTotal := checkCols(rows)
	diagTotal := checkDiags(rows)

	fmt.Println("Total:", rowTotal+colTotal+diagTotal)

	//part 2

	// splt to 3x3's
	var subgrids []string
	for i := 0; i < len(rows)-2; i++ {
		for j := 0; j < len(rows[i])-2; j++ {
			block := ""
			for k := 0; k < 3; k++ {
				block += string(rows[i][j+k])
			}
			block += "\n"
			for k := 0; k < 3; k++ {
				block += string(rows[i+1][j+k])
			}
			block += "\n"
			for k := 0; k < 3; k++ {
				block += string(rows[i+2][j+k])
			}
			block += "\n"
			subgrids = append(subgrids, block)
		}
	}
	fmt.Println(len(subgrids), "subgrids")

	//   01234567
	// 0 MMMSXXMASM
	// 1 MSAMXMSMSA
	// 2 AMXSXMAAMM
	// 3 MSAMASMSMX
	// 4 XMASAMXAMM
	// 5 XXAMMXXAMA
	// 6 SMSMSASXSS
	// 7 SAXAMASAAA
	//   MAMMMXMMMM
	//   MXMXAXMASX
	// So should be 64 subgrids of 3x3

	// checkSubGrid(subgrids[13])
	totalXs := 0
	for _, subgrid := range subgrids {
		if checkSubGrid(subgrid) {
			totalXs++
		}
	}
	fmt.Println(totalXs)
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

func checkRows(rows []string) int {
	rowFwdAndBck := 0
	for _, row := range rows {
		fwdR, _ := regexp.Compile(`XMAS`)
		bckR, _ := regexp.Compile(`SAMX`)
		found := len(fwdR.FindAllString(row, -1)) + len(bckR.FindAllString(row, -1))
		rowFwdAndBck += found
	}
	fmt.Println("Rows:", rowFwdAndBck)
	return rowFwdAndBck
}

func checkCols(rows []string) int {
	colFwdAndBck := 0
	var colString string
	rowLength := len(rows[0])

	for i := 0; i < rowLength; i++ {
		for _, row := range rows {
			colString += string(row[i])
		}
		colString += " "
	}
	cols := strings.Split(colString, " ")
	for _, col := range cols {
		fwdR, _ := regexp.Compile(`XMAS`)
		bckR, _ := regexp.Compile(`SAMX`)
		found := len(fwdR.FindAllString(col, -1)) + len(bckR.FindAllString(col, -1))
		colFwdAndBck += found
	}
	fmt.Println("Cols:", colFwdAndBck)
	return colFwdAndBck
}

func checkDiags(rows []string) int {
	rowLength := len(rows[0])
	diagString := ""

	// bottom left to top right
	for i := 0; i < 2*rowLength; i++ {
		if i < rowLength {
			diag := ""
			colPos := 0
			for j := i; j >= 0; j-- {
				diag += string(rows[j][colPos])
				colPos++
			}
			diagString += diag + " "
		} else {
			rowPos := rowLength - 1
			colPos := i - rowLength + 1
			diag := ""
			for colPos < rowLength {
				diag += string(rows[rowPos][colPos])
				rowPos--
				colPos++
			}
			diagString += diag + " "
		}
	}

	// top left to bottom right
	diagString2 := ""
	for i := 0; i < 2*rowLength; i++ {
		if i < rowLength {
			rowPos := (rowLength - 1) - i
			colPos := 0
			diag := ""
			for rowPos < rowLength {
				diag += string(rows[rowPos][colPos])
				rowPos++
				colPos++
			}
			diagString2 += diag + " "
		} else {
			rowPos := 0
			colPos := i - rowLength + 1
			diag := ""
			for colPos < rowLength {
				diag += string(rows[rowPos][colPos])
				rowPos++
				colPos++
			}
			diagString2 += diag + " "
		}
	}

	diagFwdAndBck := 0
	diags := strings.Split(diagString+" "+diagString2, " ")
	for _, diag := range diags {
		fwdR, _ := regexp.Compile(`XMAS`)
		bckR, _ := regexp.Compile(`SAMX`)
		found := len(fwdR.FindAllString(diag, -1)) + len(bckR.FindAllString(diag, -1))
		diagFwdAndBck += found
	}
	fmt.Println("Diags:", diagFwdAndBck)
	return diagFwdAndBck
}

func checkSubGrid(grid string) bool {
	// fmt.Println(grid)
	if string(grid[0]) == "M" && string(grid[2]) == "S" {
		if string(grid[5]) == "A" {
			if string(grid[8]) == "M" && string(grid[10]) == "S" {
				// fmt.Println("Is forward x-mas")
				// fmt.Println(grid)
				return true
			}
		}
	}
	if string(grid[0]) == "M" && string(grid[2]) == "M" {
		if string(grid[5]) == "A" {
			if string(grid[8]) == "S" && string(grid[10]) == "S" {
				// fmt.Println("Is forward x-mas type 2")
				// fmt.Println(grid)
				return true
			}
		}
	}
	if string(grid[0]) == "S" && string(grid[2]) == "M" {
		if string(grid[5]) == "A" {
			if string(grid[8]) == "S" && string(grid[10]) == "M" {
				// fmt.Println("Is backwards x-mas")
				// fmt.Println(grid)
				return true
			}
		}
	}
	if string(grid[0]) == "S" && string(grid[2]) == "S" {
		if string(grid[5]) == "A" {
			if string(grid[8]) == "M" && string(grid[10]) == "M" {
				// fmt.Println("Is backwards x-mas type 2")
				// fmt.Println(grid)
				return true
			}
		}
	}
	return false
}
