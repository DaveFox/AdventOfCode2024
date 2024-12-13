package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	realInput := `17639 47 3858 0 470624 9467423 5 188`
	testInput := `125 17`

	test := stringLineToNum(testInput)
	actual := stringLineToNum(realInput)

	fmt.Println("Inputs")
	fmt.Println(test, actual)
	fmt.Println(" ")

	//var workingOn []int
	//workingOn = append(workingOn, actual[0])
	workingOn := actual

	lookUp := make(map[int][]int)
	lookUp[0] = []int{1}

	for i := 0; i < 50; i++ {
		var newArray []int
		for _, val := range workingOn {
			// if new value is known increment count by lookUp table value
			newVal := applyRules(val, lookUp)
			newArray = append(newArray, newVal...)
		}
		workingOn = newArray
	}
	//fmt.Println(lookUp)

	// fmt.Println("Result:", workingOn)
	fmt.Println("Count:", len(workingOn))
	end := time.Now()
	fmt.Println("Execution time: ", end.Sub(start))

	//50 = 1344968415
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

func applyRules(in int, lookUp map[int][]int) []int {
	mapVal, ok := lookUp[in]
	if ok {
		return mapVal
	}

	if in == 0 {
		return []int{1}
	}

	count := 0
	inCopy := in
	for inCopy > 0 {
		inCopy = inCopy / 10
		count++
	}
	if math.Mod(float64(count), 2) == 0 {
		strNum := strconv.Itoa(in)
		strStart := strNum[:count/2]
		strEnd := strNum[count/2:]
		startNum, _ := strconv.Atoi(strStart)
		endNum, _ := strconv.Atoi(strEnd)
		lookUp[in] = []int{startNum, endNum}
		return []int{startNum, endNum}
	}

	newVal := in * 2024
	lookUp[in] = []int{newVal}
	return []int{newVal}
}
