package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AoC Day 5")
	inputString := fileToString()

	// testIn := `47|53
	// 97|13
	// 97|61
	// 97|47
	// 75|29
	// 61|13
	// 75|53
	// 29|13
	// 97|29
	// 53|29
	// 61|53
	// 97|53
	// 61|29
	// 47|13
	// 75|47
	// 97|75
	// 47|61
	// 75|61
	// 47|29
	// 75|13
	// 53|13
	// SPLIT
	// 75,47,61,53,29
	// 97,61,53,29,13
	// 75,29,13
	// 75,97,47,61,53
	// 61,13,29
	// 97,13,75,29,47`

	// part 1
	rules := strings.Fields(strings.Split(inputString, "\r\nSPLIT\r\n")[0])
	updates := strings.Fields(strings.Split(inputString, "\r\nSPLIT\r\n")[1]) // Real vs test \r\n vs \n

	passingUpdates := ""
	failingUpdates := ""
	for _, update := range updates {
		pass := true
		for _, rule := range rules {
			pass = applyRule(update, rule)
			if !pass {
				break
			}
		}
		if pass {
			passingUpdates += update + "  "
		} else {
			failingUpdates += update + "  "
		}
	}
	// fmt.Println("Passing updates:", passingUpdates)
	calcuateTotalOfMids(passingUpdates)

	// part 2
	// fmt.Println("Failing updates:", failingUpdates)
	failingList := strings.Fields(failingUpdates)
	nowPassing := ""
	for _, failure := range failingList {
		for i := 0; i < 20; i++ { // run multiple times. Really should run until no rule is broken but...
			for _, rule := range rules {
				failure = applyRuleAndFix(failure, rule)
			}
		}
		nowPassing += failure + "  "
	}

	// fmt.Println(nowPassing)
	calcuateTotalOfMids(nowPassing)
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

func applyRule(update string, rulePair string) bool {
	rulePages := strings.Split(rulePair, "|")
	if strings.Contains(update, rulePages[0]) && strings.Contains(update, rulePages[1]) {
		//fmt.Println("apply rule", rulePair)
		if strings.Index(update, rulePages[0]) > strings.Index(update, rulePages[1]) {
			return false
		}
	}
	return true
}

func calcuateTotalOfMids(passingUpdates string) {
	toAdd := strings.Fields(passingUpdates)
	total := 0
	for _, update := range toAdd {
		values := strings.Split(update, ",")
		mid := values[len(values)/2]
		midNum, err := strconv.Atoi(mid)
		if err != nil {
			panic(err)
		}
		total += midNum
	}
	fmt.Println("Mid totals:", total)
}

func applyRuleAndFix(update string, rulePair string) string {
	rulePages := strings.Split(rulePair, "|")
	if strings.Contains(update, rulePages[0]) && strings.Contains(update, rulePages[1]) {
		// fmt.Println("apply rule", rulePair)
		if strings.Index(update, rulePages[0]) > strings.Index(update, rulePages[1]) {
			// fmt.Println("apply fix", rulePair)
			swapper := strings.NewReplacer(rulePages[0], rulePages[1], rulePages[1], rulePages[0])
			return swapper.Replace(update)
		}
	}
	return update
}
