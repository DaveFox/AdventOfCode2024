package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("AoC Day 2")
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
