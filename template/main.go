package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Input struct{}

func parseInput(input string) Input {
	input = strings.TrimSpace(input)
	panic("parseInput not implemented")
	return Input{}
}

func part1(input Input) int {
	return 0
}

func part2(input Input) int {
	return 0
}

func main() {
	startTime := time.Now()
	input := parseInput(input)
	fmt.Println("Parsed input in", time.Since(startTime))

	startTime = time.Now()
	res1 := part1(input)
	fmt.Println("Part 1:", res1, "in", time.Since(startTime))

	startTime = time.Now()
	res2 := part2(input)
	fmt.Println("Part 2:", res2, "in", time.Since(startTime))
}
