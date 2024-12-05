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

	// 	testIn := `47|53
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

	//fmt.Println(updates[0])
	//applyRule(updates[0], rules[0])

	passingUpdates := ""
	for _, update := range updates {
		//fmt.Println("\n", update)
		pass := true
		for _, rule := range rules {
			pass = applyRule(update, rule)
			if !pass {
				break
			}
		}
		//fmt.Println(pass)
		if pass {
			passingUpdates += update + "  "
		}
	}
	// fmt.Println("")
	// fmt.Println("Passing updates:", passingUpdates)

	// for update of passingUpdate add middle vals
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
	fmt.Println(total)

	// part 2

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

func fileToString() string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	builder := new(strings.Builder)
	io.Copy(builder, file)
	return builder.String()
}
