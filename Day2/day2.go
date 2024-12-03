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
		} else {
			asNum := stringLineToNum(line)
			for index, _ := range asNum {
				subLine := removeIndex(asNum, index)
				// fmt.Println("sub", subLine)
				if isLineIncOrDec2(subLine) && diffCheck2(subLine) {
					safeCount++
					break
				}
			}
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

func stringLineToNum(stringLine string) []int {
	fields := strings.Fields(stringLine)
	var nums []int
	for _, val := range fields {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func isLineIncOrDec(line string) bool {
	increasing := true
	decreasing := true
	allLine := stringLineToNum(line)

	currI := allLine[0]
	for i := 1; i < len(allLine); i++ {
		comp := allLine[i]
		if comp < currI {
			increasing = false
			break
		}
		currI = comp
	}

	currD := allLine[0]
	for j := 1; j < len(allLine); j++ {
		comp := allLine[j]
		if comp > currD {
			decreasing = false
			break
		}
		currD = comp
	}

	return increasing || decreasing
}

func isLineIncOrDec2(line []int) bool {
	increasing := true
	decreasing := true

	currI := line[0]
	for i := 1; i < len(line); i++ {
		comp := line[i]
		if comp < currI {
			increasing = false
			break
		}
		currI = comp
	}

	currD := line[0]
	for j := 1; j < len(line); j++ {
		comp := line[j]
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
	levels := stringLineToNum(line)
	start := levels[0]

	for i := 1; i < len(levels); i++ {
		num := levels[i]
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

func diffCheck2(line []int) bool {
	levelsOk := true
	start := line[0]

	for i := 1; i < len(line); i++ {
		num := line[i]
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

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
