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

type fileBlock struct {
	start int
	len   int
	id    string
}

func main() {
	fmt.Println("AoC Day 9")
	inputString := fileToString()

	//testInput := `2333133121414131402` // p2 = 2858
	//testInput := `1333133121414131402510`
	//testInput := `2333133121414131499` // p2 = 6204
	//testInput := `48274728818` // p2 = 1752
	//testInput := `714892711` // p2 = 813
	//testInput := `12101` // p2 = 4
	//testInput := `1313165` // p2 = 169
	//testInput := `14113` // p2 = 16
	//testInput := `12345` // p2 = 132
	//testInput := `121` // p2 = 1

	// part 1
	line := strings.Fields(inputString)[0]

	blockOutput := buildBlocks(line)

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

	calculateChecksum(blockOutput)

	// part 2
	blockOutput2 := buildBlocks(line)
	// fmt.Println(blockOutput2)

	var fileAndGapSet []fileBlock
	skip := 0
	for skip < len(blockOutput2) {
		maxFileBlock := findMaxIdBlock(blockOutput2, skip)
		skip += maxFileBlock.len
		fileAndGapSet = append(fileAndGapSet, maxFileBlock)
	}
	slices.Reverse(fileAndGapSet)

	// if idblock fits then move it else check next idblock
	highToLow := slices.Clone(fileAndGapSet)
	slices.Reverse(highToLow)

	//fmt.Println(fileAndGapSet)

	for _, currHigh := range highToLow {
		for i := 0; i < len(fileAndGapSet); i++ {
			if currHigh.id == "." || currHigh.start == 0 {
				continue
			}
			if fileAndGapSet[i].id == "." && fileAndGapSet[i].len >= currHigh.len && fileAndGapSet[i].start < currHigh.start {
				//fmt.Println(currHigh.id, "fits at", fileAndGapSet[i].start)
				fileAndGapSet[i].len -= currHigh.len

				fileAndGapSet = slices.Insert(fileAndGapSet, i, currHigh)
				fileAndGapSet[i].start += len(fileAndGapSet)

				deleteIndex := slices.Index(fileAndGapSet, currHigh)
				fileAndGapSet = slices.Delete(fileAndGapSet, deleteIndex, deleteIndex+1)
				var dotBlock fileBlock
				dotBlock.id = "."
				dotBlock.len = currHigh.len
				dotBlock.start = 2 * len(fileAndGapSet)
				fileAndGapSet = slices.Insert(fileAndGapSet, deleteIndex, dotBlock)
				break
			}
		}
		for i := 0; i < len(fileAndGapSet); i++ {
			if fileAndGapSet[i].len == 0 {
				fileAndGapSet = slices.Delete(fileAndGapSet, i, i+1)
			}
		}

	}

	fmt.Println("Final")
	//fmt.Println(fileAndGapSet[0], fileAndGapSet[1], fileAndGapSet[2])
	//printBlockString(fileAndGapSet)
	var strArr []string
	for _, block := range fileAndGapSet {
		for i := 0; i < block.len; i++ {
			strArr = append(strArr, block.id)
		}
	}
	// fmt.Println(strArr)
	calculateChecksum2(strArr)
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

func buildBlocks(inputLine string) []string {
	var blockOutput []string
	fileIndex := 0
	for i := 0; i < len(inputLine); i++ {
		mod := math.Mod(float64(i), 2)
		num, err := strconv.Atoi(string(inputLine[i]))
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
	return blockOutput
}

func calculateChecksum(blocks []string) {
	checksumTotal := 0
	for i := 0; i < len(blocks); i++ {
		if string(blocks[i]) == "." {
			break
		}
		num, err := strconv.Atoi(string(blocks[i]))
		if err != nil {
			panic(err)
		}
		checksumTotal += num * i
	}
	fmt.Println("Checksum:", checksumTotal)
}

func findMaxIdBlock(blocks []string, offset int) fileBlock {
	var block fileBlock
	start := len(blocks) - 1 - offset
	block.id = blocks[start]

	for i := len(blocks) - 1 - offset; i > -2; i-- {
		if i == -1 {
			block.start = 0
			block.len = start + 1 // 99% sure this logic is right
			break
		}
		if blocks[i] != block.id {
			block.start = i + 1
			block.len = start - i
			break
		}
	}
	return block
}

func printBlockString(block []fileBlock) {
	newLine := ""
	for _, block := range block {
		for i := 0; i < block.len; i++ {
			newLine = newLine + block.id
		}
	}
	fmt.Println(newLine)
}

func calculateChecksum2(blocks []string) {
	checksumTotal := 0
	for i := 0; i < len(blocks); i++ {
		if string(blocks[i]) == "." {
			continue
		}
		num, err := strconv.Atoi(string(blocks[i]))
		if err != nil {
			panic(err)
		}
		checksumTotal += num * i
	}
	fmt.Println("Checksum:", checksumTotal)
}
