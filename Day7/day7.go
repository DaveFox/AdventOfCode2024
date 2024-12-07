package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 7")
	inputString := fileToString()

	// 	testInput := `190: 10 19
	// 3267: 81 40 27
	// 83: 17 5
	// 156: 15 6
	// 7290: 6 8 6 15
	// 161011: 16 10 13
	// 192: 17 8 14
	// 21037: 9 7 18 13
	// 292: 11 6 16 20`

	// part 1
	lines := strings.Split(inputString, "\r\n")

	total := 0
	for _, line := range lines {
		goalInt, err := strconv.Atoi(strings.Split(line, ":")[0])
		if err != nil {
			panic(err)
		}
		nums := stringLineToNum(strings.Split(line, ":")[1])

		totalAdd := addAll(nums)
		totalMult := multiplyAll(nums)

		if goalInt == totalAdd || goalInt == totalMult {
			// fmt.Println("Calibration simple add", goalInt)
			total += goalInt
			continue
		}
		total += mixed(nums, goalInt)
	}
	fmt.Println("Total part 1:", total)
	fmt.Println(" ")

	// part 2

	// debug case
	// goalInt, err := strconv.Atoi(strings.Split(lines[4], ":")[0])
	// if err != nil {
	// 	panic(err)
	// }
	// nums := stringLineToNum(strings.Split(lines[4], ":")[1])
	// mixedPart2(nums, goalInt)

	totalPart2 := 0
	for _, line := range lines {
		goalInt, err := strconv.Atoi(strings.Split(line, ":")[0])
		if err != nil {
			panic(err)
		}
		nums := stringLineToNum(strings.Split(line, ":")[1])

		totalAdd := addAll(nums)
		totalMult := multiplyAll(nums)

		if goalInt == totalAdd || goalInt == totalMult {
			// fmt.Println("Calibration simple add", goalInt)
			totalPart2 += goalInt
			continue
		}
		totalPart2 += mixedPart2(nums, goalInt)
	}
	fmt.Println("Total part 2:", totalPart2)
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

func addAll(nums []int) int {
	totalAdditive := 0
	for _, num := range nums {
		totalAdditive += num
	}
	return totalAdditive
}

func multiplyAll(nums []int) int {
	totalMuliplied := 1
	for _, num := range nums {
		totalMuliplied *= num
	}
	return totalMuliplied
}

func mixed(nums []int, goal int) int {
	depth := len(nums)
	if depth > 2 {
		var sub []int
		sub = append(sub, nums[0]+nums[1])
		sub = append(sub, nums[2:]...)
		if mixed(sub, goal) == goal {
			return goal
		}

		var sub2 []int
		sub2 = append(sub2, nums[0]*nums[1])
		sub2 = append(sub2, nums[2:]...)
		if mixed(sub2, goal) == goal {
			return goal
		}
	} else {
		if nums[0]+nums[1] == goal || nums[0]*nums[1] == goal {
			// fmt.Println("Goal", goal, "reached")
			return goal
		}
	}
	return 0
}

func mixedPart2(nums []int, goal int) int {
	depth := len(nums)
	if depth > 2 {
		var sub []int
		sub = append(sub, nums[0]+nums[1])
		sub = append(sub, nums[2:]...)
		if mixedPart2(sub, goal) == goal {
			return goal
		}

		var sub2 []int
		sub2 = append(sub2, nums[0]*nums[1])
		sub2 = append(sub2, nums[2:]...)
		if mixedPart2(sub2, goal) == goal {
			return goal
		}

		var sub3 []int
		sub3 = append(sub3, combineInts(nums[0], nums[1]))
		sub3 = append(sub3, nums[2:]...)
		if mixedPart2(sub3, goal) == goal {
			return goal
		}
	} else {
		if nums[0]+nums[1] == goal || nums[0]*nums[1] == goal || combineInts(nums[0], nums[1]) == goal {
			// fmt.Println("Goal", goal, "reached")
			return goal
		}
	}
	return 0
}

func combineInts(a int, b int) int {
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)
	combinedStr := strA + strB

	combined, err := strconv.Atoi(combinedStr)
	if err != nil {
		panic(err)
	}
	return combined
}
