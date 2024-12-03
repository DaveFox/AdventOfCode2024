package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 3")
	inputString := fileToString()

	// part 1
	// testIn := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	regCheck, _ := regexp.Compile(`mul\([\d]*,[\d]*\)`)
	instructionArray := regCheck.FindAllString(inputString, -1)
	fmt.Println("Part 1 total")
	runInstructions(instructionArray)

	//part 2
	// testIn2 := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	regCheck2, _ := regexp.Compile(`(?s)don't\(\).*?(do\(\)|$)`) // anything starting with don't() upto a do() or end of line
	filtered := regCheck2.ReplaceAllString(inputString, "")
	instructionArray2 := regCheck.FindAllString(filtered, -1)

	//fmt.Println(filtered)
	//fmt.Println(instructionArray2)
	fmt.Println("Part 2 total")
	runInstructions(instructionArray2)
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

func runInstructions(inArray []string) {
	total := 0
	for _, instruction := range inArray {
		nums := strings.Split(strings.Split(strings.Split(instruction, "mul(")[1], ")")[0], ",")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		total += (num1 * num2)
	}
	fmt.Println(total)
}
