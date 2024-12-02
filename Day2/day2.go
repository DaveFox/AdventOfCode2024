package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 2")

	fileString := fileToString()
	lines := strings.Split(fileString, "\r\n")

	//testString := "7 6 4 2 1\r\n1 2 7 8 9\r\n9 7 6 2 1\r\n1 3 2 4 5\r\n8 6 4 4 1\r\n1 3 6 7 9"
	//lines := strings.Split(testString, "\r\n")

	safeCount := 0
	for _, line := range lines {
		if isLineIncOrDec(line) && diffCheck(line) {
			safeCount++
		}
	}

	fmt.Println("Number of safe reports: ", safeCount)
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

func isLineIncOrDec(line string) bool {
	increasing := true
	decreasing := true
	allLine := strings.Fields(line)

	currI, err := strconv.Atoi(allLine[0])
	if err != nil {
		panic(err)
	}
	for i := 1; i < len(allLine); i++ {
		comp, err := strconv.Atoi(allLine[i])
		if err != nil {
			panic(err)
		}
		if comp < currI {
			increasing = false
			break
		}
		currI = comp
	}

	currD, err := strconv.Atoi(allLine[0])
	if err != nil {
		panic(err)
	}
	for j := 1; j < len(allLine); j++ {
		comp, err := strconv.Atoi(allLine[j])
		if err != nil {
			panic(err)
		}
		if comp > currD {
			decreasing = false
			break
		}
		currD = comp
	}

	return increasing || decreasing
}

func diffCheck(line string) bool {
	levelsOk := true
	levels := strings.Fields(line)
	start, err := strconv.Atoi(levels[0])
	if err != nil {
		panic(err)
	}

	for i := 1; i < len(levels); i++ {
		num, err := strconv.Atoi(levels[i])
		if err != nil {
			panic(err)
		}

		diff := start - num
		if diff < 0 {
			diff = -diff
		}
		if diff == 0 || diff > 3 {
			levelsOk = false
		}
		start = num
	}

	return levelsOk
}
