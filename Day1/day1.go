package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type leftRight struct {
	left  []int
	right []int
}

func main() {
	fmt.Println("AoC Day 1")

	asString := fileToString()
	// fmt.Println(asString)

	testIn := "3   4\r\n4   3\r\n2   5\r\n1   3\r\n3   9\r\n3   3"
	sortedListsTest := splitAndSort(testIn)
	calculateTotalDists(*sortedListsTest)
	calculateSimilarity(*sortedListsTest)

	sortedLists := splitAndSort(asString)
	calculateTotalDists(*sortedLists)
	calculateSimilarity(*sortedLists)
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

func splitAndSort(s string) *leftRight {
	var leftNums []int
	var rightNums []int

	lines := strings.Split(s, "\r\n") // Windows uses /r/n for newline....
	for i := 0; i < len(lines); i++ {
		nums := strings.Split(lines[i], "   ")

		numLeft, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		leftNums = append(leftNums, numLeft)

		numRight, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		rightNums = append(rightNums, numRight)
	}

	slices.Sort(leftNums)
	slices.Sort(rightNums)
	return &leftRight{left: leftNums, right: rightNums}
}

func calculateTotalDists(input leftRight) {
	total := 0
	for i := 0; i < len(input.left); i++ {
		diff := input.left[i] - input.right[i]

		if diff < 0 {
			diff = -diff
		}

		total += diff
	}

	fmt.Println("Total distance: ", total)
}

func calculateSimilarity(input leftRight) {
	total := 0
	for _, valLeft := range input.left {
		matches := 0
		for _, valRight := range input.right {
			if valLeft == valRight {
				matches++
			}
		}
		if matches > 0 {
			// fmt.Println("Val ", valLeft, "Matches ", matches, "Sim ", valLeft*matches)
			total += valLeft * matches
		}
	}
	fmt.Println("Total similarity: ", total)
}
