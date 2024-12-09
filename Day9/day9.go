package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 9")
	inputString := fileToString()

	//testInput := `2333133121414131402`
	//testInput := `2333133121414131402510`

	// part 1
	line := strings.Fields(inputString)[0]

	// build block slice
	var blockOutput []string
	fileIndex := 0
	for i := 0; i < len(line); i++ {
		mod := math.Mod(float64(i), 2)
		num, err := strconv.Atoi(string(line[i]))
		if err != nil {
			panic(err)
		}

		if mod == 0 {
			for num > 0 {
				blockOutput = append(blockOutput, strconv.Itoa(fileIndex))
				num--
			}
			fileIndex++
		} else {
			for num > 0 {
				blockOutput = append(blockOutput, ".")
				num--
			}
		}
	}
	// fmt.Println(blockOutput)

	// replace . with last char
	initalLength := len(blockOutput)
	for i := 0; i < initalLength; i++ {
		index := slices.Index(blockOutput, ".")
		if index == -1 || index > (len(blockOutput)-1)-i {
			break
		}

		if blockOutput[(len(blockOutput)-1)-i] != "." {
			blockOutput[index] = blockOutput[(len(blockOutput)-1)-i]
			blockOutput[(len(blockOutput)-1)-i] = "."
		}
	}
	// fmt.Println(blockOutput)

	// // calculate checksum
	checksumTotal := 0
	for i := 0; i < len(blockOutput); i++ {
		if string(blockOutput[i]) == "." {
			break
		}
		num, err := strconv.Atoi(string(blockOutput[i]))
		if err != nil {
			panic(err)
		}
		checksumTotal += num * i
	}
	fmt.Println("Checksum:", checksumTotal)

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
